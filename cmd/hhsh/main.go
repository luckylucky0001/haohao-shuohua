// hhsh 是“好好说话”的检查器：找中文里的 AI 腔调候选，核对改稿有没有丢信息。
//
//	hhsh scan  [-json] [-strict] [-skip-quoted] [文件...]                              四项零容忍扫描
//	hhsh lint  [-json] [-min 1] [-skip-quoted] [-rules FILE] [文件...]   不给文件时读标准输入
//	hhsh check [-json] [-rules FILE] 原文 改稿
//	hhsh rules [-cat 类别]
//
// 命中只是候选：同一个词在技术文档里可能是术语，在宣传稿里才是腔调。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"haohao-shuohua/internal/check"
	"haohao-shuohua/internal/lint"
	"haohao-shuohua/internal/scan"
	"haohao-shuohua/lexicon"
)

const version = "0.7.1"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "scan":
		err = cmdScan(os.Args[2:])
	case "lint":
		err = cmdLint(os.Args[2:])
	case "check":
		err = cmdCheck(os.Args[2:])
	case "rules":
		err = cmdRules(os.Args[2:])
	case "version", "-v", "--version":
		fmt.Println("hhsh", version)
	case "help", "-h", "--help":
		usage()
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "hhsh:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `hhsh：好好说话检查器

  hhsh scan  [-json] [-strict] [-skip-quoted] 文件...
                                        四项零容忍：中文里的半角标点、元话语、机械编号和量词标题、
                                        “不是 X 而是 Y”超过两处；另查标题破折号（-strict 时有半角标点就返回 1）
  hhsh lint  [-json] [-min 1] [-skip-quoted] 文件...
                                        扫描 AI 腔调候选和排版信号（不给文件读标准输入）
  hhsh check [-json] 原文 改稿          核对数字、日期、代码等是否丢失或凭空新增，统计残留
  hhsh rules [-cat 类别]                列出规则

命中只是候选，先看用途再决定改不改。
`)
}

func loadLexicon(path string) (*lexicon.Lexicon, error) {
	if path != "" {
		return lexicon.Load(path)
	}
	return lexicon.Default()
}

func readInput(path string) (string, error) {
	if path == "" || path == "-" {
		b, err := io.ReadAll(os.Stdin)
		return string(b), err
	}
	b, err := os.ReadFile(path)
	return string(b), err
}

func cmdLint(args []string) error {
	fs := flag.NewFlagSet("lint", flag.ExitOnError)
	asJSON := fs.Bool("json", false, "输出 JSON")
	minW := fs.Int("min", 1, "只显示权重不低于此值的命中（1-3）")
	rules := fs.String("rules", "", "自定义规则文件（默认用内置规则）")
	skipQuoted := fs.Bool("skip-quoted", false, "不显示引号里的命中（被引用或被讨论的说法）")
	fs.Parse(args)
	lx, err := loadLexicon(*rules)
	if err != nil {
		return err
	}
	files := fs.Args()
	if len(files) == 0 {
		files = []string{"-"}
	}
	for i, f := range files {
		text, err := readInput(f)
		if err != nil {
			return err
		}
		rep := lint.Lint(lx, text)
		if *asJSON {
			enc := json.NewEncoder(os.Stdout)
			enc.SetEscapeHTML(false)
			enc.SetIndent("", "  ")
			if err := enc.Encode(map[string]any{"file": f, "report": rep}); err != nil {
				return err
			}
			continue
		}
		if i > 0 {
			fmt.Println()
		}
		printLint(lx, f, rep, *minW, *skipQuoted)
	}
	return nil
}

func printLint(lx *lexicon.Lexicon, name string, rep lint.Report, minW int, skipQuoted bool) {
	if name == "-" {
		name = "标准输入"
	}
	fmt.Printf("%s：%d 个汉字，腔调指数 %.1f（每千字加权命中，0 最好）\n", name, rep.Stats.Han, rep.Score)
	shown := 0
	for _, h := range rep.Hits {
		if h.W < minW || (skipQuoted && h.Quoted) {
			continue
		}
		if shown == 0 {
			fmt.Println("命中：")
		}
		shown++
		keep := ""
		if h.Keep != "" {
			keep = "（保留：" + h.Keep + "）"
		}
		q := ""
		if h.Quoted {
			q = "（在引号里，可能是被引用的说法）"
		}
		fmt.Printf("  第%d行 [%s·%d] “%s”%s → %s%s\n      %s\n", h.Line, h.CatName, h.W, h.Match, q, h.Hint, keep, h.Context)
	}
	if shown == 0 {
		fmt.Println("没有命中规则。")
	}
	printStats(rep.Stats)
	fmt.Println("提示：命中只是候选，先判断它在这里是不是术语、引语或作者本来的意思。")
}

func printStats(st lint.Stats) {
	var sig []string
	add := func(n int, label string) {
		if n > 0 {
			sig = append(sig, fmt.Sprintf("%s %d", label, n))
		}
	}
	add(st.Emoji, "表情")
	add(st.Bold, "加粗")
	add(st.Headings, "标题")
	add(st.ListLines, "列表行")
	add(st.Dashes, "破折号")
	add(st.Exclaims, "感叹号")
	add(st.HalfPunct, "半角标点")
	add(st.FourCharRuns, "四字格连用")
	add(st.LongSentences, "超长句")
	add(st.DeHeavy, "“的”字堆叠句")
	if len(sig) > 0 {
		fmt.Println("排版信号：" + strings.Join(sig, " · "))
	}
	for _, p := range st.Parallels {
		fmt.Printf("排比：%s\n", p)
	}
	if st.CV >= 0 {
		note := ""
		if st.CV < 0.35 {
			note = "（偏低：句子长短太均匀）"
		}
		fmt.Printf("节奏：%d 句，平均 %.0f 字，句长变异系数 %.2f%s\n", st.Sentences, st.MeanLen, st.CV, note)
	}
	if st.Han >= 200 && st.ConnPer1k > 25 {
		fmt.Printf("连接词：每千字 %.0f 个，偏多\n", st.ConnPer1k)
	}
}

func cmdCheck(args []string) error {
	fs := flag.NewFlagSet("check", flag.ExitOnError)
	asJSON := fs.Bool("json", false, "输出 JSON")
	rules := fs.String("rules", "", "自定义规则文件（默认用内置规则）")
	fs.Parse(args)
	if fs.NArg() != 2 {
		return fmt.Errorf("用法：hhsh check 原文 改稿")
	}
	lx, err := loadLexicon(*rules)
	if err != nil {
		return err
	}
	orig, err := readInput(fs.Arg(0))
	if err != nil {
		return err
	}
	rw, err := readInput(fs.Arg(1))
	if err != nil {
		return err
	}
	res := check.Compare(lx, orig, rw)
	if *asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		return enc.Encode(res)
	}
	fmt.Println("保护项核对（数字、日期、时间、版本、代码、链接、书名、英文名）：")
	if len(res.Missing) == 0 {
		fmt.Println("  没有丢失。")
	} else {
		var parts []string
		for _, t := range res.Missing {
			parts = append(parts, fmt.Sprintf("%s“%s”", t.Kind, t.Text))
		}
		fmt.Printf("  可能丢失 %d 项：%s\n", len(res.Missing), strings.Join(parts, "，"))
		fmt.Println("  （中文数字改写成阿拉伯数字也会出现在这里，逐项确认）")
	}
	if len(res.Added) > 0 {
		var parts []string
		for _, t := range res.Added {
			parts = append(parts, fmt.Sprintf("%s“%s”", t.Kind, t.Text))
		}
		fmt.Printf("  原文没有、改稿新增 %d 项：%s（确认不是编出来的）\n", len(res.Added), strings.Join(parts, "，"))
	}
	ratio := 0.0
	if res.OrigHan > 0 {
		ratio = float64(res.NewHan) / float64(res.OrigHan) * 100
	}
	fmt.Printf("改动：编辑距离占 %.0f%%；篇幅 %d → %d 个汉字（%.0f%%）\n", res.EditRatio*100, res.OrigHan, res.NewHan, ratio)
	fmt.Printf("腔调指数：%.1f → %.1f；原文命中 %d 处，改稿残留 %d 处，新增 %d 处\n", res.OrigScore, res.NewScore, res.OrigHits, len(res.Residual), len(res.Introduced))
	for _, h := range res.Residual {
		fmt.Printf("  残留 第%d行 [%s] “%s” → %s\n", h.Line, h.CatName, h.Match, h.Hint)
	}
	for _, h := range res.Introduced {
		fmt.Printf("  新增 第%d行 [%s] “%s” → %s\n", h.Line, h.CatName, h.Match, h.Hint)
	}
	return nil
}

func cmdRules(args []string) error {
	fs := flag.NewFlagSet("rules", flag.ExitOnError)
	cat := fs.String("cat", "", "只列这个类别（英文键名，如 jargon）")
	fs.Parse(args)
	lx, err := lexicon.Default()
	if err != nil {
		return err
	}
	fmt.Printf("规则表 %s，共 %d 条\n", lx.Version, len(lx.Rules))
	for _, k := range lx.CatKeys() {
		if *cat != "" && *cat != k {
			continue
		}
		fmt.Printf("\n## %s（%s）\n", lx.CatName(k), k)
		for _, r := range lx.Rules {
			if r.Cat != k {
				continue
			}
			keep := ""
			if r.Keep != "" {
				keep = "；保留：" + r.Keep
			}
			fmt.Printf("  [%d] %-16s %s → %s%s\n", r.W, r.ID, r.Re, r.Hint, keep)
		}
	}
	return nil
}

func cmdScan(args []string) error {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	asJSON := fs.Bool("json", false, "输出 JSON")
	strict := fs.Bool("strict", false, "中文正文里有半角标点时返回非零退出码")
	skipQuoted := fs.Bool("skip-quoted", false, "不报中文引号里的命中（被引用或被讨论的说法）")
	fs.Parse(args)
	files := fs.Args()
	if len(files) == 0 {
		files = []string{"-"}
	}
	dirty := false
	for i, f := range files {
		text, err := readInput(f)
		if err != nil {
			return err
		}
		rep := scan.ScanWith(text, scan.Options{SkipQuoted: *skipQuoted})
		if len(rep.HalfWidth) > 0 {
			dirty = true
		}
		if *asJSON {
			enc := json.NewEncoder(os.Stdout)
			enc.SetEscapeHTML(false)
			enc.SetIndent("", "  ")
			if err := enc.Encode(map[string]any{"file": f, "report": rep, "clean": rep.Clean()}); err != nil {
				return err
			}
			continue
		}
		if i > 0 {
			fmt.Println()
		}
		printScan(f, rep)
	}
	if *strict && dirty {
		os.Exit(1)
	}
	return nil
}

func printScan(name string, rep scan.Report) {
	if name == "-" {
		name = "标准输入"
	}
	fmt.Printf("%s：\n", name)
	section := func(title, rule string, hits []scan.Hit, bad bool) {
		mark := "通过"
		if bad {
			mark = "要处理"
		}
		fmt.Printf("  %s：%d 处（%s，%s）\n", title, len(hits), rule, mark)
		for _, h := range hits {
			fmt.Printf("    第%d行 [%s] “%s”  %s\n", h.Line, h.Kind, h.Match, h.Context)
		}
	}
	section("中文里的半角标点", "必须清零，代码、英文、千分位、网址除外", rep.HalfWidth, len(rep.HalfWidth) > 0)
	section("元话语", "删掉或换成具体判断", rep.Meta, len(rep.Meta) > 0)
	section("机械编号和量词标题", "文章里改成能立住的小标题；PRD、技术方案等查阅文档的编号可保留", rep.Numbering, len(rep.Numbering) > 0)
	section("“不是 X 而是 Y”及变体", fmt.Sprintf("全文最多 %d 处", rep.NotButMax), rep.NotButIs, rep.NotButOver())
	section("标题里的破折号", "标题不用破折号", rep.TitleDash, len(rep.TitleDash) > 0)
	fmt.Println("  半角标点能判死，改到零；其余几项只看字面，是候选，结合上下文判断。")
}
