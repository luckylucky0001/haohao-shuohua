package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
)

// bootstrapWin 按用例重抽样，返回某方案胜率的 95% 区间。
// 每次从全部用例里有放回地抽同样多的题，重算胜场/局数。
func bootstrapWin(cond string, ids []string, perCase map[string]map[string][2]int, iters int, rng *rand.Rand) (lo, hi float64) {
	var have []string
	for _, id := range ids {
		if pc, ok := perCase[id][cond]; ok && pc[1] > 0 {
			have = append(have, id)
		}
	}
	if len(have) == 0 {
		return 0, 0
	}
	rates := make([]float64, iters)
	for i := range rates {
		w, g := 0, 0
		for range have {
			pc := perCase[have[rng.Intn(len(have))]][cond]
			w += pc[0]
			g += pc[1]
		}
		rates[i] = float64(w) / float64(g)
	}
	sort.Float64s(rates)
	return rates[int(0.025*float64(iters))], rates[int(0.975*float64(iters))-1]
}

// binomTwoSided 是符号检验：n 次里赢 k 次，假设双方各一半时的双侧 p 值。
func binomTwoSided(k, n int) float64 {
	if n == 0 {
		return 1
	}
	if k > n-k {
		k = n - k
	}
	p := 0.0
	for i := 0; i <= k; i++ {
		p += math.Exp(lgammaChoose(n, i) - float64(n)*math.Ln2)
	}
	p *= 2
	if p > 1 {
		p = 1
	}
	return p
}

func lgammaChoose(n, k int) float64 {
	a, _ := math.Lgamma(float64(n + 1))
	b, _ := math.Lgamma(float64(k + 1))
	c, _ := math.Lgamma(float64(n - k + 1))
	return a - b - c
}

func renderStats(conds []string, perCase map[string]map[string][2]int, ahead map[string]map[string]map[string]bool) string {
	var ids []string
	for id := range perCase {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	if len(ids) < 2 {
		return ""
	}
	rng := rand.New(rand.NewSource(1))
	var b strings.Builder
	b.WriteString("\n## 胜率的 95% 区间\n\n按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。\n\n| 方案 | 胜率 | 95% 区间 |\n|---|---:|---|\n")
	for _, c := range conds {
		w, g := 0, 0
		for _, id := range ids {
			pc := perCase[id][c]
			w += pc[0]
			g += pc[1]
		}
		if g == 0 {
			continue
		}
		lo, hi := bootstrapWin(c, ids, perCase, 2000, rng)
		fmt.Fprintf(&b, "| %s | %.0f%% | %.0f%%–%.0f%% |\n", c, float64(w)/float64(g)*100, lo*100, hi*100)
	}
	b.WriteString("\n## 两两对比\n\n逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。\n\n| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |\n|---|---|---:|---:|---:|\n")
	for i := 0; i < len(conds); i++ {
		for j := i + 1; j < len(conds); j++ {
			a, c := conds[i], conds[j]
			wa, wb := 0, 0
			for _, id := range ids {
				if ahead[id][a][c] {
					wa++
				} else if ahead[id][c][a] {
					wb++
				}
			}
			if wa+wb == 0 {
				continue
			}
			fmt.Fprintf(&b, "| %s | %s | %d | %d | %.3f |\n", a, c, wa, wb, binomTwoSided(wa, wa+wb))
		}
	}
	return b.String()
}
