// Package lexicon 提供内置的中文 AI 腔调规则表（zh.json）。
//
// 规则只用来找“候选”，不用来下结论：同一个词在技术文档里可能是术语，
// 在宣传稿里才是腔调。每条规则的 keep 字段写的就是应该保留的情形。
package lexicon

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
)

//go:embed zh.json
var builtin []byte

// Rule 是一条规则。
type Rule struct {
	ID   string `json:"id"`
	Cat  string `json:"cat"`
	W    int    `json:"w"`
	Re   string `json:"re"`
	Hint string `json:"hint"`
	Keep string `json:"keep,omitempty"`

	rx *regexp.Regexp
}

// Regexp 返回编译后的正则。
func (r *Rule) Regexp() *regexp.Regexp { return r.rx }

// Lexicon 是一份规则表。
type Lexicon struct {
	Version    string            `json:"version"`
	Note       string            `json:"note"`
	Categories map[string]string `json:"categories"`
	Rules      []*Rule           `json:"rules"`
}

// CatName 返回类别的中文名。
func (l *Lexicon) CatName(cat string) string {
	if n, ok := l.Categories[cat]; ok {
		return n
	}
	return cat
}

// CatKeys 按固定顺序返回类别。
func (l *Lexicon) CatKeys() []string {
	order := []string{"trace", "opener", "closer", "chat", "jargon", "hype", "lofty", "template", "transl", "literary", "webslang", "engineer", "service"}
	seen := map[string]bool{}
	var out []string
	for _, k := range order {
		if _, ok := l.Categories[k]; ok {
			out = append(out, k)
			seen[k] = true
		}
	}
	var rest []string
	for k := range l.Categories {
		if !seen[k] {
			rest = append(rest, k)
		}
	}
	sort.Strings(rest)
	return append(out, rest...)
}

// Default 返回内置规则表。
func Default() (*Lexicon, error) { return Parse(builtin) }

// MustDefault 返回内置规则表，出错时 panic（内置表由测试保证可用）。
func MustDefault() *Lexicon {
	l, err := Default()
	if err != nil {
		panic(err)
	}
	return l
}

// Load 从文件读取规则表，用于试验自定义规则。
func Load(path string) (*Lexicon, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(b)
}

// Parse 解析并编译规则表。
func Parse(b []byte) (*Lexicon, error) {
	var l Lexicon
	if err := json.Unmarshal(b, &l); err != nil {
		return nil, err
	}
	ids := map[string]bool{}
	for _, r := range l.Rules {
		if r.ID == "" || r.Re == "" {
			return nil, fmt.Errorf("规则缺少 id 或 re：%+v", r)
		}
		if ids[r.ID] {
			return nil, fmt.Errorf("规则 id 重复：%s", r.ID)
		}
		ids[r.ID] = true
		if _, ok := l.Categories[r.Cat]; !ok {
			return nil, fmt.Errorf("规则 %s 的类别 %q 未定义", r.ID, r.Cat)
		}
		if r.W < 1 || r.W > 3 {
			return nil, fmt.Errorf("规则 %s 的权重应为 1-3", r.ID)
		}
		rx, err := regexp.Compile(r.Re)
		if err != nil {
			return nil, fmt.Errorf("规则 %s 正则错误：%w", r.ID, err)
		}
		r.rx = rx
	}
	return &l, nil
}
