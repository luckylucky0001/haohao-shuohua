// hheval 是 haohao-shuohua 的评测辅助工具：
//
//	hheval inputs -cases evals/cases -kinds rewrite,control -out tasks.md
//	    生成交给改写模型的任务单（只含请求与原文，不含评分要点）。
//	hheval blind  -cases evals/cases -runs DIR -conds a,b,c -out DIR -seed 1
//	    把多个方案的输出打乱成匿名的判分包，映射另存 mapping.json。
//	hheval report -cases evals/cases -runs DIR -judge DIR -out report.md
//	    汇总判分结果和检查器指标。
//	hheval notes  -judge DIR -cond 方案名
//	    列出判官对某个方案逐题的打分与简评。
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"haohao-shuohua/internal/evalcase"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "inputs":
		err = cmdInputs(os.Args[2:])
	case "blind":
		err = cmdBlind(os.Args[2:])
	case "report":
		err = cmdReport(os.Args[2:])
	case "notes":
		err = cmdNotes(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "hheval:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "用法: hheval inputs|blind|report|notes [参数]")
}

func splitList(s string) []string {
	var out []string
	for _, p := range strings.Split(s, ",") {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

func filterCases(cases []evalcase.Case, kinds []string) []evalcase.Case {
	if len(kinds) == 0 {
		return cases
	}
	want := map[string]bool{}
	for _, k := range kinds {
		want[k] = true
	}
	var out []evalcase.Case
	for _, c := range cases {
		if want[c.Kind] || want[c.ID] {
			out = append(out, c)
		}
	}
	return out
}

func cmdInputs(args []string) error {
	fs := flag.NewFlagSet("inputs", flag.ExitOnError)
	casesDir := fs.String("cases", "evals/cases", "用例目录")
	kinds := fs.String("kinds", "", "只要这些 kind 或 ID，逗号分隔；为空表示全部")
	out := fs.String("out", "", "输出文件，默认标准输出")
	fs.Parse(args)

	cases, err := evalcase.Load(*casesDir)
	if err != nil {
		return err
	}
	cases = filterCases(cases, splitList(*kinds))
	var b strings.Builder
	b.WriteString("# 任务单\n\n每个任务都是一次独立的新对话：用户只发了下面这条消息（请求 + 贴进来的材料）。材料在 ~~~~ 围栏里，围栏本身不属于材料。\n\n")
	for _, c := range cases {
		fmt.Fprintf(&b, "## 任务 %s\n\n用户消息：\n\n%s\n\n", c.ID, c.Request)
		if c.Source != "" {
			fmt.Fprintf(&b, "~~~~text\n%s\n~~~~\n\n", c.Source)
		}
	}
	if *out == "" {
		fmt.Print(b.String())
		return nil
	}
	return os.WriteFile(*out, []byte(b.String()), 0o644)
}
