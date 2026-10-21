// Package lint 扫描中文文本里的 AI 腔调候选，并统计排版与节奏信号。
package lint

import (
	"math"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"haohao-shuohua/internal/textutil"
	"haohao-shuohua/lexicon"
)

// Hit 是一处规则命中。
type Hit struct {
	RuleID  string `json:"rule"`
	Cat     string `json:"cat"`
	CatName string `json:"cat_name"`
	W       int    `json:"w"`
	Match   string `json:"match"`
	Line    int    `json:"line"`
	Context string `json:"context"`
	Hint    string `json:"hint"`
	Keep    string `json:"keep,omitempty"`
	Quoted  bool   `json:"quoted,omitempty"` // 在引号里：可能是被讨论或被引用的说法
	Start   int    `json:"-"`
	End     int    `json:"-"`
}

// Stats 是排版与节奏信号。只报数，阈值仅供参考。
type Stats struct {
	Han           int      `json:"han"`
	Sentences     int      `json:"sentences"`
	MeanLen       float64  `json:"mean_sentence_han"`
	CV            float64  `json:"sentence_len_cv"` // 句数不足 8 时为 -1
	LongSentences int      `json:"long_sentences"`  // 超过 60 个汉字的句子
	DeHeavy       int      `json:"de_heavy"`        // 一句里“的”出现 4 次及以上
	Emoji         int      `json:"emoji"`
	Bold          int      `json:"bold"`
	Headings      int      `json:"headings"`
	ListLines     int      `json:"list_lines"`
	Dashes        int      `json:"dashes"`
	Exclaims      int      `json:"exclaims"`
	HalfPunct     int      `json:"half_width_punct"`
	FourCharRuns  int      `json:"four_char_runs"` // 三个及以上四字格连用
	Parallels     []string `json:"parallels"`      // 三连排比
	Connectives   int      `json:"connectives"`
	ConnPer1k     float64  `json:"connectives_per_1k"`
}

// Report 是一次扫描的结果。
type Report struct {
	Hits  []Hit          `json:"hits"`
	Score float64        `json:"score"` // 每千汉字的加权命中
	ByCat map[string]int `json:"by_cat"`
	Stats Stats          `json:"stats"`
}

// Lint 扫描文本。代码块、行内代码和链接不参与规则匹配。
func Lint(lx *lexicon.Lexicon, text string) Report {
	masked := textutil.MaskCode(text)
	type cand struct {
		r          *lexicon.Rule
		start, end int
	}
	var cands []cand
	for _, r := range lx.Rules {
		target := masked
		if r.Cat == "trace" {
			// 工具痕迹常藏在链接里，要对原文匹配。
			target = text
		}
		for _, loc := range r.Regexp().FindAllStringIndex(target, -1) {
			if loc[1] > loc[0] {
				cands = append(cands, cand{r, loc[0], loc[1]})
			}
		}
	}
	// 重叠时保留权重高、匹配长的那条。
	sort.Slice(cands, func(i, j int) bool {
		a, b := cands[i], cands[j]
		if a.r.W != b.r.W {
			return a.r.W > b.r.W
		}
		if (a.end - a.start) != (b.end - b.start) {
			return (a.end - a.start) > (b.end - b.start)
		}
		return a.start < b.start
	})
	var taken [][2]int
	overlaps := func(s, e int) bool {
		for _, t := range taken {
			if s < t[1] && e > t[0] {
				return true
			}
		}
		return false
	}
	rep := Report{ByCat: map[string]int{}}
	weight := 0
	for _, c := range cands {
		if overlaps(c.start, c.end) {
			continue
		}
		taken = append(taken, [2]int{c.start, c.end})
		h := Hit{
			RuleID: c.r.ID, Cat: c.r.Cat, CatName: lx.CatName(c.r.Cat), W: c.r.W,
			Match:   text[c.start:c.end],
			Line:    textutil.LineOf(text, c.start),
			Context: textutil.Context(text, c.start, c.end, 12),
			Hint:    c.r.Hint, Keep: c.r.Keep, Start: c.start, End: c.end,
			Quoted: inQuotes(text, c.start),
		}
		rep.Hits = append(rep.Hits, h)
		rep.ByCat[c.r.Cat]++
		weight += c.r.W
	}
	sort.Slice(rep.Hits, func(i, j int) bool { return rep.Hits[i].Start < rep.Hits[j].Start })
	rep.Stats = computeStats(text, masked)
	if rep.Stats.Han > 0 {
		rep.Score = float64(weight) * 1000 / float64(rep.Stats.Han)
	}
	return rep
}

var (
	boldRe     = regexp.MustCompile(`\*\*[^*\n]+\*\*|__[^_\n]+__`)
	headingRe  = regexp.MustCompile(`(?m)^\s*(#{1,6}\s|[一二三四五六七八九十]+、|（[一二三四五六七八九十]+）|【[^】\n]{1,20}】)`)
	listRe     = regexp.MustCompile(`(?m)^\s*([-*+•·]\s|\d+[.、)）]\s?|[①②③④⑤⑥⑦⑧⑨⑩])`)
	dashRe     = regexp.MustCompile(`—+|--`)
	halfRe     = regexp.MustCompile(`\p{Han}[,;:!?]`)
	sentSep    = regexp.MustCompile(`[。！？!?；;\n]`)
	fourSep    = regexp.MustCompile(`[、，,]`)
	clauseSep  = regexp.MustCompile(`[，,；;、：:]`)
	connectRe  = regexp.MustCompile(`此外|另外|然而|因此|所以|同时|而且|并且|从而|进而|首先|其次|最后|另一方面|与此同时|不仅|而是|但是|可是|不过|此外`)
	exclaimSet = "！!"
)

func isEmoji(r rune) bool {
	switch {
	case r >= 0x1F300 && r <= 0x1FAFF:
		return true
	case r >= 0x2600 && r <= 0x27BF:
		return true
	case r >= 0x1F000 && r <= 0x1F2FF:
		return true
	case r == 0x2B50 || r == 0x2B55 || r == 0x203C || r == 0x2049:
		return true
	}
	return false
}

func computeStats(text, masked string) Stats {
	var st Stats
	st.Han = textutil.HanCount(masked)
	for _, r := range masked {
		if isEmoji(r) {
			st.Emoji++
		}
		if strings.ContainsRune(exclaimSet, r) {
			st.Exclaims++
		}
	}
	st.Bold = len(boldRe.FindAllString(masked, -1))
	st.Headings = len(headingRe.FindAllString(masked, -1))
	st.ListLines = len(listRe.FindAllString(masked, -1))
	st.Dashes = len(dashRe.FindAllString(masked, -1))
	st.HalfPunct = len(halfRe.FindAllString(masked, -1))
	st.FourCharRuns = fourCharRuns(masked)
	conns := connectRe.FindAllString(masked, -1)
	st.Connectives = len(conns)
	if st.Han > 0 {
		st.ConnPer1k = float64(st.Connectives) * 1000 / float64(st.Han)
	}

	var lens []float64
	for _, s := range textutil.Sentences(masked) {
		n := textutil.HanCount(s)
		if n >= 4 {
			lens = append(lens, float64(n))
			if n > 60 {
				st.LongSentences++
			}
			if strings.Count(s, "的") >= 4 {
				st.DeHeavy++
			}
		}
		if p := parallel(s); p != "" {
			st.Parallels = append(st.Parallels, p)
		}
	}
	st.Sentences = len(lens)
	st.CV = -1
	if len(lens) > 0 {
		sum := 0.0
		for _, v := range lens {
			sum += v
		}
		st.MeanLen = sum / float64(len(lens))
		if len(lens) >= 8 && st.MeanLen > 0 {
			v := 0.0
			for _, x := range lens {
				v += (x - st.MeanLen) * (x - st.MeanLen)
			}
			st.CV = math.Sqrt(v/float64(len(lens))) / st.MeanLen
		}
	}
	return st
}

// parallel 找一句里连续三个以上分句用同样两个字开头的排比，例如
// “是每一位……，是每一个……，是每一次……”“忙着开会，忙着回消息，忙着刷……”。
func parallel(sentence string) string {
	parts := clauseSep.Split(sentence, -1)
	run, best := 1, 1
	bestEnd := -1
	prefix := func(s string) string {
		s = strings.TrimSpace(s)
		var b strings.Builder
		n := 0
		for _, r := range s {
			if !unicode.Is(unicode.Han, r) {
				return ""
			}
			b.WriteRune(r)
			n++
			if n == 2 {
				return b.String()
			}
		}
		return ""
	}
	for i := 1; i < len(parts); i++ {
		p0, p1 := prefix(parts[i-1]), prefix(parts[i])
		if p0 != "" && p0 == p1 {
			run++
			if run > best {
				best, bestEnd = run, i
			}
		} else {
			run = 1
		}
	}
	if best >= 3 {
		start := bestEnd - best + 1
		s := strings.Join(parts[start:bestEnd+1], "，")
		if utf8.RuneCountInString(s) > 40 {
			rs := []rune(s)
			s = string(rs[:40]) + "…"
		}
		return s
	}
	return ""
}

// fourCharRuns 数“三个及以上四字短语连用”，例如“高效便捷、安全可靠、灵活多样”。
// 只认被顿号或逗号隔开、恰好四个汉字的分句，避免把长句切出假阳性。
func fourCharRuns(text string) int {
	n := 0
	for _, sent := range sentSep.Split(text, -1) {
		run := 0
		for _, c := range fourSep.Split(sent, -1) {
			c = strings.TrimSpace(c)
			if textutil.HanCount(c) == 4 && utf8.RuneCountInString(c) == 4 {
				run++
				if run == 3 {
					n++
				}
			} else {
				run = 0
			}
		}
	}
	return n
}

// inQuotes 判断偏移是否落在同一行的引号（“” 或 「」 或 『』）之内。
func inQuotes(text string, off int) bool {
	lineStart := strings.LastIndex(text[:off], "\n") + 1
	seg := text[lineStart:off]
	pairs := [][2]string{{"“", "”"}, {"「", "」"}, {"『", "』"}}
	for _, p := range pairs {
		if strings.Count(seg, p[0]) > strings.Count(seg, p[1]) {
			return true
		}
	}
	return false
}
