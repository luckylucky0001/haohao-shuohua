package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"haohao-shuohua/internal/evalcase"
)

// readOutput 读取某个方案在某条用例上的输出：正文 + 可选的附言。
func readOutput(runs, cond, id string) (text, notes string, ok bool) {
	b, err := os.ReadFile(filepath.Join(runs, cond, id+".md"))
	if err != nil {
		return "", "", false
	}
	n, _ := os.ReadFile(filepath.Join(runs, cond, id+".notes.md"))
	return strings.TrimSpace(string(b)), strings.TrimSpace(string(n)), true
}

func seedFor(seed int64, id string) int64 {
	h := fnv.New64a()
	h.Write([]byte(id))
	return seed ^ int64(h.Sum64())
}

const labels = "ABCDEFGH"

func cmdBlind(args []string) error {
	fs := flag.NewFlagSet("blind", flag.ExitOnError)
	casesDir := fs.String("cases", "evals/cases", "用例目录")
	runs := fs.String("runs", "", "本轮输出目录（下面每个方案一个子目录）")
	conds := fs.String("conds", "", "参与比较的方案，逗号分隔")
	kinds := fs.String("kinds", "", "只要这些 kind 或 ID")
	batches := fs.Int("batches", 2, "拆成几个判分包")
	seed := fs.Int64("seed", 1, "随机种子")
	out := fs.String("out", "", "判分包输出目录")
	fs.Parse(args)
	if *runs == "" || *out == "" || *conds == "" {
		return fmt.Errorf("需要 -runs、-conds、-out")
	}
	cases, err := evalcase.Load(*casesDir)
	if err != nil {
		return err
	}
	cases = filterCases(cases, splitList(*kinds))
	condList := splitList(*conds)
	if err := os.MkdirAll(*out, 0o755); err != nil {
		return err
	}
	mapping := map[string]map[string]string{}
	bundles := make([]strings.Builder, *batches)
	for i := range bundles {
		fmt.Fprintf(&bundles[i], "# 判分包 %d\n\n", i+1)
	}
	for ci, c := range cases {
		var present []string
		for _, cond := range condList {
			if _, _, ok := readOutput(*runs, cond, c.ID); ok {
				present = append(present, cond)
			}
		}
		if len(present) < 2 {
			continue
		}
		rng := rand.New(rand.NewSource(seedFor(*seed, c.ID)))
		rng.Shuffle(len(present), func(i, j int) { present[i], present[j] = present[j], present[i] })
		m := map[string]string{}
		b := &bundles[ci%*batches]
		writeCaseHeader(b, c)
		for i, cond := range present {
			label := string(labels[i])
			m[label] = cond
			text, notes, _ := readOutput(*runs, cond, c.ID)
			fmt.Fprintf(b, "### 候选 %s\n\n~~~~text\n%s\n~~~~\n\n", label, text)
			if notes != "" {
				fmt.Fprintf(b, "候选 %s 在正文之外对用户说的话：\n\n~~~~text\n%s\n~~~~\n\n", label, notes)
			}
		}
		mapping[c.ID] = m
	}
	for i := range bundles {
		p := filepath.Join(*out, fmt.Sprintf("bundle-%d.md", i+1))
		if err := os.WriteFile(p, []byte(bundles[i].String()), 0o644); err != nil {
			return err
		}
	}
	mb, _ := json.MarshalIndent(mapping, "", "  ")
	if err := os.WriteFile(filepath.Join(*out, "mapping.json"), mb, 0o644); err != nil {
		return err
	}
	var ids []string
	for id := range mapping {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	fmt.Printf("写出 %d 个判分包，%d 条用例：%s\n", *batches, len(ids), strings.Join(ids, " "))
	return nil
}

func writeCaseHeader(b *strings.Builder, c evalcase.Case) {
	fmt.Fprintf(b, "## 用例 %s（%s）\n\n", c.ID, kindName(c.Kind))
	fmt.Fprintf(b, "- 场景：%s\n- 语体期待：%s\n- 用户请求：%s\n", c.Scene, c.Register, c.Request)
	list := func(title string, items []string) {
		if len(items) == 0 {
			return
		}
		fmt.Fprintf(b, "- %s：\n", title)
		for _, it := range items {
			fmt.Fprintf(b, "  - %s\n", it)
		}
	}
	list("必须保留的信息", c.MustKeep)
	list("陷阱（不该改坏的地方）", c.Traps)
	list("应清掉的腔调（参考，不是硬性清单）", c.Targets)
	b.WriteString("\n")
	if c.Source != "" {
		fmt.Fprintf(b, "### 原文\n\n~~~~text\n%s\n~~~~\n\n", c.Source)
	}
}

func kindName(k string) string {
	switch k {
	case "rewrite":
		return "改稿"
	case "control":
		return "对照：原文本来就自然"
	case "write":
		return "从零写作"
	}
	return k
}
