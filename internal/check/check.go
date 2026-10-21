// Package check 对照原文和改稿：保护项有没有丢、有没有凭空多出来，
// 改了多少，腔调残留还有多少。
package check

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"haohao-shuohua/internal/lint"
	"haohao-shuohua/internal/textutil"
	"haohao-shuohua/lexicon"
)

// Token 是一个应当原样保留的片段。
type Token struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

type pattern struct {
	kind string
	re   *regexp.Regexp
}

// 顺序即优先级：先匹配到的长片段占位，后面的短片段不再重复计。
var patterns = []pattern{
	{"代码", regexp.MustCompile("`[^`\\n]+`")},
	{"链接", regexp.MustCompile(`https?://[^\s)>\]]+`)},
	{"书名", regexp.MustCompile(`《[^》\n]{1,30}》`)},
	{"日期", regexp.MustCompile(`\d{4}\s*年\s*\d{1,2}\s*月(?:\s*\d{1,2}\s*日)?|\d{1,2}\s*月\s*\d{1,2}\s*日|\d{4}-\d{1,2}-\d{1,2}`)},
	{"时间", regexp.MustCompile(`\d{1,2}\s*[:：]\s*\d{2}`)},
	{"版本", regexp.MustCompile(`[vV]?\d+(?:\.\d+){2,3}`)},
	{"数字", regexp.MustCompile(`\d+(?:\.\d+)?\s*(?:%|％|ms|毫秒|秒|分钟|小时|个工作日|工作日|天|周|个月|月|年|倍|次|人|台|个|条|万|亿|元|块|GB|MB|KB|TB|行|通|套|名|页|号)?`)},
	{"中文数字", regexp.MustCompile(`[一二三四五六七八九十两半百千]+(?:个)?(?:天|周|个月|年|小时|分钟|点|次|倍|人|下午|上午|晚上|件事|条)`)},
	{"英文", regexp.MustCompile(`[A-Za-z][A-Za-z0-9_./+-]*[A-Za-z0-9]|[A-Za-z]`)},
}

var (
	spaceRe   = regexp.MustCompile(`\s+`)
	percentRe = regexp.MustCompile(`百分之([零〇一二两三四五六七八九十百千]+)`)
	cnNumRe   = regexp.MustCompile(`[零〇一二两三四五六七八九十百千万亿]+`)
	wanRe     = regexp.MustCompile(`(\d+(?:\.\d+)?)万`)
	geRe      = regexp.MustCompile(`(\d)个`)
)

// normalize 抹平写法差异，只为比较用：去空白、统一全半角、
// 把中文数字和“1.3 万”这类写法都换成阿拉伯数字。
func normalize(s string) string {
	s = spaceRe.ReplaceAllString(s, "")
	r := strings.NewReplacer("％", "%", "：", ":", "，", ",", "（", "(", "）", ")", "～", "~", "—", "-", "–", "-")
	s = r.Replace(s)
	s = wanRe.ReplaceAllStringFunc(s, func(m string) string {
		f, err := strconv.ParseFloat(strings.TrimSuffix(m, "万"), 64)
		if err != nil {
			return m
		}
		v := f * 10000
		if v == float64(int64(v)) {
			return strconv.FormatInt(int64(v), 10)
		}
		return m
	})
	s = percentRe.ReplaceAllStringFunc(s, func(m string) string {
		return strconv.FormatInt(cnToInt(strings.TrimPrefix(m, "百分之")), 10) + "%"
	})
	s = cnNumRe.ReplaceAllStringFunc(s, func(m string) string {
		return strconv.FormatInt(cnToInt(m), 10)
	})
	s = geRe.ReplaceAllString(s, "$1")
	return s
}

// cnToInt 把“九十五”“一万三千”“两”这类中文数字转成整数。
func cnToInt(s string) int64 {
	digit := map[rune]int64{'零': 0, '〇': 0, '一': 1, '二': 2, '两': 2, '三': 3, '四': 4, '五': 5, '六': 6, '七': 7, '八': 8, '九': 9}
	unit := map[rune]int64{'十': 10, '百': 100, '千': 1000}
	big := map[rune]int64{'万': 10000, '亿': 100000000}
	var total, section, number int64
	for _, r := range s {
		if d, ok := digit[r]; ok {
			number = d
		} else if u, ok := unit[r]; ok {
			if number == 0 {
				number = 1
			}
			section += number * u
			number = 0
		} else if b, ok := big[r]; ok {
			section += number
			if section == 0 {
				section = 1
			}
			total += section * b
			section, number = 0, 0
		}
	}
	return total + section + number
}

// Protected 提取文本里的保护项（去重，保持出现顺序）。
func Protected(text string) []Token {
	type span struct{ s, e int }
	var taken []span
	var out []Token
	seen := map[string]bool{}
	for _, p := range patterns {
		for _, loc := range p.re.FindAllStringIndex(text, -1) {
			s, e := loc[0], loc[1]
			ok := true
			for _, t := range taken {
				if s < t.e && e > t.s {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			taken = append(taken, span{s, e})
			t := strings.TrimSpace(text[s:e])
			if t == "" {
				continue
			}
			// “一次”“一点”“一天”多是虚指，不当保护项。
			if p.kind == "中文数字" && strings.HasPrefix(t, "一") && !strings.ContainsAny(t, "十百千万") {
				continue
			}
			key := p.kind + "|" + normalize(t)
			if seen[key] {
				continue
			}
			seen[key] = true
			out = append(out, Token{p.kind, t})
		}
	}
	return out
}

// Result 是对照结果。
type Result struct {
	Missing    []Token    `json:"missing"`
	Added      []Token    `json:"added"`
	EditRatio  float64    `json:"edit_ratio"`
	OrigHan    int        `json:"orig_han"`
	NewHan     int        `json:"new_han"`
	OrigScore  float64    `json:"orig_score"`
	NewScore   float64    `json:"new_score"`
	OrigHits   int        `json:"orig_hits"`
	Residual   []lint.Hit `json:"residual"`   // 原文就有、改稿还在
	Introduced []lint.Hit `json:"introduced"` // 原文没有、改稿新增
	NewStats   lint.Stats `json:"new_stats"`
	OrigStats  lint.Stats `json:"orig_stats"`
}

// Compare 对照原文与改稿。
func Compare(lx *lexicon.Lexicon, orig, rewrite string) Result {
	var r Result
	no, nr := normalize(orig), normalize(rewrite)
	for _, t := range Protected(orig) {
		if !strings.Contains(nr, normalize(t.Text)) {
			r.Missing = append(r.Missing, t)
		}
	}
	for _, t := range Protected(rewrite) {
		if t.Kind == "英文" && len(t.Text) < 2 {
			continue
		}
		if !strings.Contains(no, normalize(t.Text)) {
			r.Added = append(r.Added, t)
		}
	}
	r.EditRatio = textutil.EditRatio(orig, rewrite)
	lo, ln := lint.Lint(lx, orig), lint.Lint(lx, rewrite)
	r.OrigHan, r.NewHan = lo.Stats.Han, ln.Stats.Han
	r.OrigScore, r.NewScore = lo.Score, ln.Score
	r.OrigHits = len(lo.Hits)
	r.OrigStats, r.NewStats = lo.Stats, ln.Stats
	inOrig := map[string]int{}
	for _, h := range lo.Hits {
		inOrig[h.RuleID+"|"+h.Match]++
	}
	for _, h := range ln.Hits {
		k := h.RuleID + "|" + h.Match
		if inOrig[k] > 0 {
			inOrig[k]--
			r.Residual = append(r.Residual, h)
		} else {
			r.Introduced = append(r.Introduced, h)
		}
	}
	sort.SliceStable(r.Missing, func(i, j int) bool { return kindRank(r.Missing[i].Kind) < kindRank(r.Missing[j].Kind) })
	return r
}

func kindRank(k string) int {
	for i, p := range patterns {
		if p.kind == k {
			return i
		}
	}
	return len(patterns)
}
