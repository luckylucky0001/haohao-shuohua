// Package textutil 放中文文本处理的小工具。
package textutil

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// HanCount 统计汉字个数。
func HanCount(s string) int {
	n := 0
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			n++
		}
	}
	return n
}

var (
	fenceRe  = regexp.MustCompile("(?s)(```|~~~).*?(```|~~~)")
	inlineRe = regexp.MustCompile("`[^`\\n]+`")
	urlRe    = regexp.MustCompile(`https?://[^\s)>\]]+`)
)

// MaskCode 把代码块、行内代码和链接替换成等长的空白（byte 数不变，保留换行），
// 这样规则不会命中代码里的内容，命中位置的 byte 偏移和行号也不变。
func MaskCode(s string) string {
	blank := func(m string) string {
		b := []byte(m)
		for i, c := range b {
			if c != '\n' {
				b[i] = ' '
			}
		}
		return string(b)
	}
	s = fenceRe.ReplaceAllStringFunc(s, blank)
	s = inlineRe.ReplaceAllStringFunc(s, blank)
	s = urlRe.ReplaceAllStringFunc(s, blank)
	return s
}

// LineOf 返回 byte 偏移所在的行号（从 1 开始）。
func LineOf(s string, off int) int {
	if off > len(s) {
		off = len(s)
	}
	return strings.Count(s[:off], "\n") + 1
}

// Context 返回偏移附近的一小段文字，用于展示命中位置。
func Context(s string, start, end, radius int) string {
	// 往前找 radius 个字符
	i := start
	for k := 0; k < radius && i > 0; k++ {
		_, size := utf8.DecodeLastRuneInString(s[:i])
		if s[i-size] == '\n' {
			break
		}
		i -= size
	}
	j := end
	for k := 0; k < radius && j < len(s); k++ {
		r, size := utf8.DecodeRuneInString(s[j:])
		if r == '\n' {
			break
		}
		j += size
	}
	return strings.TrimSpace(s[i:j])
}

var sentEndRe = regexp.MustCompile(`[^。！？!?\n]+[。！？!?]+|[^。！？!?\n]+$`)

// Sentences 粗切句子（按句末标点和换行），返回每句文字。
func Sentences(s string) []string {
	var out []string
	for _, line := range strings.Split(s, "\n") {
		for _, m := range sentEndRe.FindAllString(line, -1) {
			if strings.TrimSpace(m) != "" {
				out = append(out, strings.TrimSpace(m))
			}
		}
	}
	return out
}

// Levenshtein 计算两段文字按字符（rune）的编辑距离。
func Levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	if len(ra) == 0 {
		return len(rb)
	}
	if len(rb) == 0 {
		return len(ra)
	}
	prev := make([]int, len(rb)+1)
	cur := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min3(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev, cur = cur, prev
	}
	return prev[len(rb)]
}

func min3(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

// EditRatio 返回编辑距离占较长文本长度的比例（0 表示一字未改）。
func EditRatio(a, b string) float64 {
	a, b = strings.TrimSpace(a), strings.TrimSpace(b)
	la, lb := utf8.RuneCountInString(a), utf8.RuneCountInString(b)
	m := la
	if lb > m {
		m = lb
	}
	if m == 0 {
		return 0
	}
	return float64(Levenshtein(a, b)) / float64(m)
}
