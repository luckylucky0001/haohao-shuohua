// Package scan 做四项零容忍扫描：中文正文里的半角标点、元话语、
// 机械编号和量词标题、“不是 X 而是 Y”及其变体超过两处。另外检查标题里的破折号。
//
// 半角标点能判得死，应当改到零；其余几项只看字面，判不了意图，
// 结果只作候选，由人决定改不改。
package scan

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"haohao-shuohua/internal/textutil"
)

// Hit 是一处命中。
type Hit struct {
	Kind    string `json:"kind"`
	Line    int    `json:"line"`
	Match   string `json:"match"`
	Context string `json:"context"`
	off     int
}

// Report 是一份扫描结果。
type Report struct {
	HalfWidth []Hit `json:"half_width"`
	Meta      []Hit `json:"meta"`
	Numbering []Hit `json:"numbering"`
	NotButIs  []Hit `json:"not_but_is"`
	TitleDash []Hit `json:"title_dash"`
	NotButMax int   `json:"not_but_is_max"`
}

// NotButOver 报告“不是 X 而是 Y”变体是否超过上限。
func (r Report) NotButOver() bool { return len(r.NotButIs) > r.NotButMax }

// Clean 表示四项都过：半角标点为零、元话语为零、没有机械编号、变体不超过上限、标题没有破折号。
func (r Report) Clean() bool {
	return len(r.HalfWidth) == 0 && len(r.Meta) == 0 && len(r.Numbering) == 0 && !r.NotButOver() && len(r.TitleDash) == 0
}

var (
	metaRe = regexp.MustCompile(`值得(深思|关注|玩味|注意的是|一提的是|强调的是)|不难(看出|发现|理解)|不禁让人(思考|感叹|深思)|引人深思|发人深省|耐人寻味`)

	notButRes = []*regexp.Regexp{
		regexp.MustCompile(`不是[^，。！？；\n]{1,30}[，,]\s*(而是|而|是)`),
		regexp.MustCompile(`不是[^，。！？；\n]{1,20}[，,][^，。！？；\n]{1,12}才是`),
		regexp.MustCompile(`不在于?[^，。！？；\n]{1,20}[，,]\s*而在于?`),
		regexp.MustCompile(`看似[^，。！？；\n]{1,30}[，,]\s*(实则|其实|实际上)`),
		regexp.MustCompile(`你以为[^。！？；\n]{1,30}[，,]\s*(其实|实际上)`),
		regexp.MustCompile(`真正[^，。！？；\n]{1,12}的(是|，)`),
		regexp.MustCompile(`回头(才|再)(发现|看)`),
		regexp.MustCompile(`不重要[，,]\s*重要的是`),
		regexp.MustCompile(`管的是[^，。！？；\n]{1,12}[，,]\s*不管`),
	}

	ordinalLineRe = regexp.MustCompile(`^(#{1,6}\s*)?([一二三四五六七八九十]+、|第[一二三四五六七八九十0-9]+[章节部分篇]|其[一二三四五六七八九十][，、：:\s]|[0-9]+[.、]\s*\S)`)
	quantRe       = regexp.MustCompile(`[一二三四五六七八九十两0-9]+\s*(板斧|把刀|大法宝|大支柱|大抓手|个维度|大维度|重境界|把钥匙|步走|字诀|招|刀)|第[一二三四五六七八九十0-9]+(刀|招|式|板斧)`)
	headingRe     = regexp.MustCompile(`^#{1,6}\s`)
	cnOrdinalRe   = regexp.MustCompile(`^([一二三四五六七八九十]+、|第[一二三四五六七八九十0-9]+[章节部分篇])`)
)

// Options 控制扫描行为。
type Options struct {
	// SkipQuoted 跳过中文引号“”里的命中（被引用或被讨论的说法）。半角标点不受影响。
	SkipQuoted bool
}

// Scan 扫描一段文字。代码块、行内代码和链接不参与扫描。
func Scan(text string) Report { return ScanWith(text, Options{}) }

// ScanWith 按选项扫描。
func ScanWith(text string, opt Options) Report {
	rep := scanAll(text)
	if opt.SkipQuoted {
		rep.Meta = dropQuoted(text, rep.Meta)
		rep.NotButIs = dropQuoted(text, rep.NotButIs)
		rep.Numbering = dropQuoted(text, rep.Numbering)
	}
	return rep
}

func scanAll(text string) Report {
	masked := textutil.MaskCode(text)
	rep := Report{NotButMax: 2}
	rep.HalfWidth = halfWidth(masked, text)
	for _, loc := range metaRe.FindAllStringIndex(masked, -1) {
		rep.Meta = append(rep.Meta, mkHit("元话语", text, loc[0], loc[1]))
	}
	for _, re := range notButRes {
		for _, loc := range re.FindAllStringIndex(masked, -1) {
			rep.NotButIs = append(rep.NotButIs, mkHit("不是X而是Y", text, loc[0], loc[1]))
		}
	}
	sortHits(rep.NotButIs)
	off := 0
	firstHeading := true
	for _, line := range strings.SplitAfter(masked, "\n") {
		trim := strings.TrimSpace(line)
		isHeading := headingRe.MatchString(trim)
		if trim != "" && (isHeading || isShortTitleLine(trim)) {
			re := ordinalLineRe
			if !isHeading {
				re = cnOrdinalRe // 没有 # 的短行，只认中文序号，避免把列表当标题
			}
			if loc := re.FindStringIndex(trim); loc != nil {
				start := off + strings.Index(line, trim)
				rep.Numbering = append(rep.Numbering, mkHit("机械编号", text, start+loc[0], start+loc[1]))
			}
			if isHeading && firstHeading {
				firstHeading = false
				if i := strings.Index(trim, "——"); i >= 0 {
					start := off + strings.Index(line, trim)
					rep.TitleDash = append(rep.TitleDash, mkHit("标题破折号", text, start+i, start+i+len("——")))
				}
			}
		}
		if trim != "" {
			start := off + strings.Index(line, trim)
			if isHeading {
				for _, loc := range quantRe.FindAllStringIndex(trim, -1) {
					rep.Numbering = append(rep.Numbering, mkHit("量词标题", text, start+loc[0], start+loc[1]))
				}
			} else {
				// 正文里只抓“第一刀”这类明显的量词标签。
				for _, loc := range quantRe.FindAllStringIndex(trim, -1) {
					m := trim[loc[0]:loc[1]]
					if strings.HasPrefix(m, "第") {
						rep.Numbering = append(rep.Numbering, mkHit("量词标签", text, start+loc[0], start+loc[1]))
					}
				}
			}
		}
		off += len(line)
	}
	return rep
}

// isShortTitleLine 粗判一行是不是没有 # 的小标题：很短、不以句末标点结尾。
func isShortTitleLine(s string) bool {
	n := utf8.RuneCountInString(s)
	if n == 0 || n > 20 {
		return false
	}
	r, _ := utf8.DecodeLastRuneInString(s)
	return !strings.ContainsRune("。！？；：，、.!?;:,", r)
}

// halfWidth 找夹在中文里的半角逗号、句号、分号、冒号。
// 判据：紧挨着的前一个字或后一个字是汉字或全角标点。数字千分位、小数、时间、
// 文件名、英文句子里的标点两边都不是汉字，不会命中。
func halfWidth(masked, text string) []Hit {
	var hits []Hit
	for i := 0; i < len(masked); {
		r, size := utf8.DecodeRuneInString(masked[i:])
		if r == ',' || r == '.' || r == ';' || r == ':' {
			prev, _ := utf8.DecodeLastRuneInString(masked[:i])
			next, _ := utf8.DecodeRuneInString(masked[i+size:])
			if isCJK(prev) || isCJK(next) {
				// 省略号 "..." 只报一次
				j := i + size
				for j < len(masked) && masked[j] == '.' && r == '.' {
					j++
				}
				hits = append(hits, mkHit("半角标点", text, i, j))
				i = j
				continue
			}
		}
		i += size
	}
	return hits
}

func isCJK(r rune) bool {
	if unicode.Is(unicode.Han, r) {
		return true
	}
	// 全角标点和中文引号
	return strings.ContainsRune("，。、；：？！“”‘’（）《》【】……—", r)
}

func mkHit(kind, text string, s, e int) Hit {
	return Hit{Kind: kind, Line: textutil.LineOf(text, s), Match: text[s:e], Context: textutil.Context(text, s, e, 14), off: s}
}

// dropQuoted 去掉落在中文引号里的命中。
func dropQuoted(text string, hits []Hit) []Hit {
	var out []Hit
	for _, h := range hits {
		if !inQuotes(text, h.off) {
			out = append(out, h)
		}
	}
	return out
}

// inQuotes 判断偏移是否在同一行的“”之间。
func inQuotes(text string, off int) bool {
	lineStart := strings.LastIndex(text[:off], "\n") + 1
	before := text[lineStart:off]
	return strings.Count(before, "“") > strings.Count(before, "”")
}

func sortHits(h []Hit) {
	for i := 1; i < len(h); i++ {
		for j := i; j > 0 && h[j].Line < h[j-1].Line; j-- {
			h[j], h[j-1] = h[j-1], h[j]
		}
	}
}
