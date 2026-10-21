package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"haohao-shuohua/internal/check"
	"haohao-shuohua/internal/evalcase"
	"haohao-shuohua/internal/lint"
	"haohao-shuohua/lexicon"
)

type candScore struct {
	Fidelity string   `json:"fidelity"`
	Lost     []string `json:"lost"`
	Added    []string `json:"added"`
	Deai     float64  `json:"deai"`
	Natural  float64  `json:"natural"`
	Register float64  `json:"register"`
	Overedit float64  `json:"overedit"`
	Tics     []string `json:"tics"`
	Note     string   `json:"note"`
}

type caseResult struct {
	ID         string               `json:"id"`
	Candidates map[string]candScore `json:"candidates"`
	Ranking    []string             `json:"ranking"`
}

type judgeFile struct {
	Cases []caseResult `json:"cases"`
}

type agg struct {
	n                       int
	pass, minor, fail       int
	deai, natural, register float64
	quality                 float64
	overeditN               int
	overedit                float64
	rankNorm                float64
	wins, games             int
	outScore                float64
	outScoreN               int
	residual, introduced    int
	missing, added          int
	ctrlEdit                float64
	ctrlN                   int
	lenRatio                float64
	lenN                    int
	byKindQ                 map[string][2]float64 // kind -> (sum quality, n)
	byKindR                 map[string][2]float64 // kind -> (sum rankNorm, n)
}

func newAgg() *agg {
	return &agg{byKindQ: map[string][2]float64{}, byKindR: map[string][2]float64{}}
}

func cmdReport(args []string) error {
	fs := flag.NewFlagSet("report", flag.ExitOnError)
	casesDir := fs.String("cases", "evals/cases", "用例目录")
	runs := fs.String("runs", "", "本轮输出目录")
	judgeDir := fs.String("judge", "", "判分目录（mapping.json 和 result-*.json）")
	out := fs.String("out", "", "报告输出文件，默认标准输出")
	title := fs.String("title", "评测报告", "报告标题")
	order := fs.String("order", "", "方案在表里的顺序，逗号分隔")
	svgOut := fs.String("svg", "", "另外输出一张胜率与质量分的 SVG 图")
	labelsFlag := fs.String("labels", "", "图里方案的显示名，形如 haohao=好好说话,baseline=不装 skill")
	highlight := fs.String("highlight", "", "图里高亮的方案，逗号分隔")
	skip := fs.String("skip", "", "不计入报告的用例 ID，逗号分隔（例如只看新用例时跳过旧用例）")
	fs.Parse(args)
	skipSet := map[string]bool{}
	for _, id := range splitList(*skip) {
		skipSet[id] = true
	}
	if *runs == "" || *judgeDir == "" {
		return fmt.Errorf("需要 -runs 和 -judge")
	}
	cases, err := evalcase.Load(*casesDir)
	if err != nil {
		return err
	}
	byID := map[string]evalcase.Case{}
	for _, c := range cases {
		byID[c.ID] = c
	}
	var mapping map[string]map[string]string
	mb, err := os.ReadFile(filepath.Join(*judgeDir, "mapping.json"))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(mb, &mapping); err != nil {
		return err
	}
	files, _ := filepath.Glob(filepath.Join(*judgeDir, "result-*.json"))
	results := map[string]caseResult{}
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		var jf judgeFile
		if err := json.Unmarshal(b, &jf); err != nil {
			return fmt.Errorf("%s: %w", f, err)
		}
		for _, cr := range jf.Cases {
			if !skipSet[cr.ID] {
				results[cr.ID] = cr
			}
		}
	}
	lx := lexicon.MustDefault()
	aggs := map[string]*agg{}
	get := func(cond string) *agg {
		if aggs[cond] == nil {
			aggs[cond] = newAgg()
		}
		return aggs[cond]
	}
	var ids []string
	for id := range mapping {
		if !skipSet[id] {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)

	var fidelityLines, ticLines, rankLines, missingIDs []string
	// 逐题记录：每个方案在这题的胜场和局数，以及两两之间谁排在前面，用于置信区间和配对比较。
	perCase := map[string]map[string][2]int{}
	ahead := map[string]map[string]map[string]bool{} // id -> 赢家 -> 输家
	for _, id := range ids {
		c := byID[id]
		cr, ok := results[id]
		if !ok {
			missingIDs = append(missingIDs, id)
			continue
		}
		m := mapping[id]
		k := len(m)
		pos := map[string]int{}
		for i, label := range cr.Ranking {
			pos[label] = i
		}
		var rankNames []string
		for _, label := range cr.Ranking {
			rankNames = append(rankNames, m[label])
		}
		rankLines = append(rankLines, fmt.Sprintf("| %s | %s | %s |", id, kindName(c.Kind), strings.Join(rankNames, " > ")))
		perCase[id] = map[string][2]int{}
		ahead[id] = map[string]map[string]bool{}
		for label, cond := range m {
			a := get(cond)
			s := cr.Candidates[label]
			a.n++
			switch s.Fidelity {
			case "pass":
				a.pass++
			case "minor":
				a.minor++
			default:
				a.fail++
			}
			if s.Fidelity != "pass" {
				fidelityLines = append(fidelityLines, fmt.Sprintf("| %s | %s | %s | %s | %s |", id, cond, s.Fidelity, strings.Join(s.Lost, "；"), strings.Join(s.Added, "；")))
			}
			if len(s.Tics) > 0 {
				ticLines = append(ticLines, fmt.Sprintf("| %s | %s | %s |", id, cond, strings.Join(s.Tics, "；")))
			}
			q := (s.Deai + s.Natural + s.Register) / 3
			a.deai += s.Deai
			a.natural += s.Natural
			a.register += s.Register
			a.quality += q
			kq := a.byKindQ[c.Kind]
			a.byKindQ[c.Kind] = [2]float64{kq[0] + q, kq[1] + 1}
			if c.Kind == "control" {
				a.overeditN++
				a.overedit += s.Overedit
			}
			p, ok := pos[label]
			if ok && k > 1 {
				rn := float64(p) / float64(k-1)
				a.rankNorm += rn
				kr := a.byKindR[c.Kind]
				a.byKindR[c.Kind] = [2]float64{kr[0] + rn, kr[1] + 1}
				for other := range m {
					if other == label {
						continue
					}
					po, ok2 := pos[other]
					if !ok2 {
						continue
					}
					a.games++
					pc := perCase[id][cond]
					pc[1]++
					if p < po {
						a.wins++
						pc[0]++
						if ahead[id][cond] == nil {
							ahead[id][cond] = map[string]bool{}
						}
						ahead[id][cond][m[other]] = true
					}
					perCase[id][cond] = pc
				}
			}
			text, _, _ := readOutput(*runs, cond, id)
			if c.Source != "" {
				res := check.Compare(lx, c.Source, text)
				a.outScore += res.NewScore
				a.outScoreN++
				a.residual += len(res.Residual)
				a.introduced += len(res.Introduced)
				a.missing += len(res.Missing)
				a.added += len(res.Added)
				if res.OrigHan > 0 {
					a.lenRatio += float64(res.NewHan) / float64(res.OrigHan)
					a.lenN++
				}
				if c.Kind == "control" {
					a.ctrlEdit += res.EditRatio
					a.ctrlN++
				}
			} else {
				rep := lint.Lint(lx, text)
				a.outScore += rep.Score
				a.outScoreN++
				a.introduced += len(rep.Hits)
			}
		}
	}

	var conds []string
	if *order != "" {
		for _, c := range splitList(*order) {
			if aggs[c] != nil {
				conds = append(conds, c)
			}
		}
	}
	for c := range aggs {
		found := false
		for _, x := range conds {
			if x == c {
				found = true
			}
		}
		if !found {
			conds = append(conds, c)
		}
	}

	if *svgOut != "" {
		names := map[string]string{}
		for _, kv := range splitList(*labelsFlag) {
			if k, v, ok := strings.Cut(kv, "="); ok {
				names[k] = v
			}
		}
		hl := map[string]bool{}
		for _, h := range splitList(*highlight) {
			hl[h] = true
		}
		if err := os.WriteFile(*svgOut, []byte(renderSVG(conds, aggs, names, hl)), 0o644); err != nil {
			return err
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n", *title)
	fmt.Fprintf(&b, "判分 %d 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。\n\n", len(results))
	if len(missingIDs) > 0 {
		fmt.Fprintf(&b, "缺判分：%s\n\n", strings.Join(missingIDs, " "))
	}
	b.WriteString("## 总表\n\n| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |\n|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, cond := range conds {
		a := aggs[cond]
		n := float64(a.n)
		over := "-"
		if a.overeditN > 0 {
			over = fmt.Sprintf("%.2f", a.overedit/float64(a.overeditN))
		}
		win := 0.0
		if a.games > 0 {
			win = float64(a.wins) / float64(a.games)
		}
		os := 0.0
		if a.outScoreN > 0 {
			os = a.outScore / float64(a.outScoreN)
		}
		fmt.Fprintf(&b, "| %s | %d | %d/%d/%d | %.2f | %.2f | %.2f | **%.2f** | %.2f | %.0f%% | %s | %.1f | %d | %d | %d | %d |\n",
			cond, a.n, a.pass, a.minor, a.fail, a.deai/n, a.natural/n, a.register/n, a.quality/n, a.rankNorm/n, win*100, over, os, a.residual, a.introduced, a.missing, a.added)
	}
	b.WriteString(renderStats(conds, perCase, ahead))
	b.WriteString("\n“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。\n\n")

	b.WriteString("## 按用例类型\n\n| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |\n|---|---|---|---|---:|---:|\n")
	for _, cond := range conds {
		a := aggs[cond]
		cell := func(k string) string {
			q, r := a.byKindQ[k], a.byKindR[k]
			if q[1] == 0 {
				return "-"
			}
			rr := "-"
			if r[1] > 0 {
				rr = fmt.Sprintf("%.2f", r[0]/r[1])
			}
			return fmt.Sprintf("%.2f / %s", q[0]/q[1], rr)
		}
		ce := "-"
		if a.ctrlN > 0 {
			ce = fmt.Sprintf("%.0f%%", a.ctrlEdit/float64(a.ctrlN)*100)
		}
		lr := "-"
		if a.lenN > 0 {
			lr = fmt.Sprintf("%.0f%%", a.lenRatio/float64(a.lenN)*100)
		}
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %s |\n", cond, cell("rewrite"), cell("control"), cell("write"), ce, lr)
	}

	b.WriteString("\n## 逐题名次（从好到差）\n\n| 用例 | 类型 | 名次 |\n|---|---|---|\n")
	b.WriteString(strings.Join(rankLines, "\n") + "\n")

	b.WriteString("\n## 保真问题\n\n")
	if len(fidelityLines) == 0 {
		b.WriteString("没有。\n")
	} else {
		sort.Strings(fidelityLines)
		b.WriteString("| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |\n|---|---|---|---|---|\n")
		b.WriteString(strings.Join(fidelityLines, "\n") + "\n")
	}
	b.WriteString("\n## 新长出来的腔调\n\n")
	if len(ticLines) == 0 {
		b.WriteString("没有。\n")
	} else {
		sort.Strings(ticLines)
		b.WriteString("| 用例 | 方案 | 判官指出的新腔调 |\n|---|---|---|\n")
		b.WriteString(strings.Join(ticLines, "\n") + "\n")
	}
	if *out == "" {
		fmt.Print(b.String())
		return nil
	}
	return os.WriteFile(*out, []byte(b.String()), 0o644)
}

// cmdNotes 列出判官对某个方案每道题的打分和简评，便于逐题分析。
func cmdNotes(args []string) error {
	fs := flag.NewFlagSet("notes", flag.ExitOnError)
	judgeDir := fs.String("judge", "", "判分目录")
	cond := fs.String("cond", "", "方案名")
	fs.Parse(args)
	var mapping map[string]map[string]string
	mb, err := os.ReadFile(filepath.Join(*judgeDir, "mapping.json"))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(mb, &mapping); err != nil {
		return err
	}
	files, _ := filepath.Glob(filepath.Join(*judgeDir, "result-*.json"))
	var lines []string
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		var jf judgeFile
		if err := json.Unmarshal(b, &jf); err != nil {
			return err
		}
		for _, cr := range jf.Cases {
			for label, s := range cr.Candidates {
				if mapping[cr.ID][label] != *cond {
					continue
				}
				rank := 0
				for i, l := range cr.Ranking {
					if l == label {
						rank = i + 1
					}
				}
				lines = append(lines, fmt.Sprintf("%s 第%d/%d 保真%s 去味%.0f 自然%.0f 语体%.0f 过改%.0f｜%s｜丢:%s｜增:%s｜腔:%s",
					cr.ID, rank, len(cr.Ranking), s.Fidelity, s.Deai, s.Natural, s.Register, s.Overedit, s.Note,
					strings.Join(s.Lost, "；"), strings.Join(s.Added, "；"), strings.Join(s.Tics, "；")))
			}
		}
	}
	sort.Strings(lines)
	fmt.Println(strings.Join(lines, "\n"))
	return nil
}

// renderSVG 画两栏横向条形图：左边胜率，右边质量分，下方标保真失败数。
func renderSVG(conds []string, aggs map[string]*agg, names map[string]string, hl map[string]bool) string {
	const (
		rowH   = 34
		top    = 46
		labelW = 150
		barW   = 230
		gap    = 60
	)
	h := top + rowH*len(conds) + 30
	w := labelW + barW*2 + gap + 130
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" font-family="Noto Sans CJK SC, sans-serif">`, w, h, w, h)
	fmt.Fprintf(&b, `<text x="%d" y="22" font-size="13" fill="#44505f">胜率（两两比较中排在对方前面的比例）</text>`, labelW)
	fmt.Fprintf(&b, `<text x="%d" y="22" font-size="13" fill="#44505f">质量分（去味·自然·语体，满分 5）</text>`, labelW+barW+gap)
	for i, c := range conds {
		a := aggs[c]
		y := top + i*rowH
		name := c
		if n, ok := names[c]; ok {
			name = n
		}
		color := "#b9b2a5"
		weight := "400"
		if hl[c] {
			color = "#1d2a3a"
			weight = "700"
		}
		fmt.Fprintf(&b, `<text x="%d" y="%d" font-size="13" text-anchor="end" fill="#1d2a3a" font-weight="%s">%s</text>`, labelW-10, y+17, weight, name)
		win := 0.0
		if a.games > 0 {
			win = float64(a.wins) / float64(a.games)
		}
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%d" height="20" fill="#efebe3"/>`, labelW, y+3, barW)
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%.1f" height="20" fill="%s"/>`, labelW, y+3, win*barW, color)
		fmt.Fprintf(&b, `<text x="%.1f" y="%d" font-size="12" fill="#1d2a3a">%.0f%%</text>`, float64(labelW)+win*barW+6, y+18, win*100)
		q := 0.0
		if a.n > 0 {
			q = a.quality / float64(a.n)
		}
		x2 := labelW + barW + gap
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%d" height="20" fill="#efebe3"/>`, x2, y+3, barW)
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%.1f" height="20" fill="%s"/>`, x2, y+3, q/5*barW, color)
		fail := ""
		if a.fail > 0 {
			fail = fmt.Sprintf("  保真失败 %d", a.fail)
		}
		fmt.Fprintf(&b, `<text x="%.1f" y="%d" font-size="12" fill="#1d2a3a">%.2f<tspan fill="#b8432f">%s</tspan></text>`, float64(x2)+q/5*barW+6, y+18, q, fail)
	}
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-size="11" fill="#8a8378">题数：`, labelW, h-8)
	for i, c := range conds {
		name := c
		if n, ok := names[c]; ok {
			name = n
		}
		sep := "，"
		if i == 0 {
			sep = ""
		}
		fmt.Fprintf(&b, "%s%s %d", sep, name, aggs[c].n)
	}
	b.WriteString(`</text></svg>`)
	return b.String()
}
