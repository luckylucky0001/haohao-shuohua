// Package evalcase 读取 evals/cases 下的评测用例。
//
// 用例文件是带简单头信息的 Markdown：
//
//	---
//	id: R01
//	kind: rewrite | control | write
//	scene: ...
//	register: ...
//	request: ...
//	must_keep:
//	  - ...
//	targets:
//	  - ...
//	traps:
//	  - ...
//	---
//	原文（write 类用例为空）
package evalcase

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Case 是一条评测用例。
type Case struct {
	ID       string
	Kind     string // rewrite / control / write
	Scene    string
	Register string
	Request  string
	MustKeep []string
	Targets  []string
	Traps    []string
	Source   string
	Path     string
}

// Load 读取目录下全部 .md 用例，按 ID 排序。
func Load(dir string) ([]Case, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return nil, err
	}
	var out []Case
	for _, f := range files {
		c, err := ParseFile(f)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// ParseFile 解析单个用例文件。
func ParseFile(path string) (Case, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Case{}, err
	}
	c, err := Parse(string(b))
	c.Path = path
	return c, err
}

// Parse 解析用例文本。
func Parse(s string) (Case, error) {
	var c Case
	s = strings.ReplaceAll(s, "\r\n", "\n")
	if !strings.HasPrefix(s, "---\n") {
		return c, fmt.Errorf("缺少头信息")
	}
	rest := s[4:]
	end := strings.Index(rest, "\n---\n")
	if end < 0 {
		// 允许头信息后没有正文
		if strings.HasSuffix(rest, "\n---") {
			end = len(rest) - 4
		} else {
			return c, fmt.Errorf("头信息没有结束标记")
		}
	}
	head := rest[:end]
	body := ""
	if end+5 <= len(rest) {
		body = rest[end+5:]
	}
	c.Source = strings.TrimSpace(body)

	var curList *[]string
	sc := bufio.NewScanner(strings.NewReader(head))
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.HasPrefix(line, "  - ") || strings.HasPrefix(line, "- ") {
			if curList == nil {
				return c, fmt.Errorf("列表项不属于任何字段：%q", line)
			}
			item := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "- "))
			*curList = append(*curList, item)
			continue
		}
		k, v, ok := strings.Cut(line, ":")
		if !ok {
			return c, fmt.Errorf("无法解析：%q", line)
		}
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		curList = nil
		switch k {
		case "id":
			c.ID = v
		case "kind":
			c.Kind = v
		case "scene":
			c.Scene = v
		case "register":
			c.Register = v
		case "request":
			c.Request = v
		case "must_keep":
			curList = &c.MustKeep
		case "targets":
			curList = &c.Targets
		case "traps":
			curList = &c.Traps
		default:
			return c, fmt.Errorf("未知字段 %q", k)
		}
	}
	if c.ID == "" || c.Kind == "" || c.Request == "" {
		return c, fmt.Errorf("id/kind/request 不能为空")
	}
	switch c.Kind {
	case "rewrite", "control":
		if c.Source == "" {
			return c, fmt.Errorf("%s 类用例需要原文", c.Kind)
		}
	case "write":
	default:
		return c, fmt.Errorf("未知 kind %q", c.Kind)
	}
	return c, sc.Err()
}
