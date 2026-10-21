package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestBinomTwoSided(t *testing.T) {
	// 10 局赢 9 局：双侧 p = 2*(1+10)/1024 ≈ 0.0215
	if p := binomTwoSided(9, 10); math.Abs(p-0.021484375) > 1e-9 {
		t.Fatalf("got %v", p)
	}
	if p := binomTwoSided(5, 10); p != 1 {
		t.Fatalf("even split should be 1, got %v", p)
	}
}

func TestBootstrapWinBounds(t *testing.T) {
	pc := map[string]map[string][2]int{}
	var ids []string
	for i := 0; i < 40; i++ {
		id := string(rune('A'+i%26)) + string(rune('a'+i/26))
		ids = append(ids, id)
		w := 0
		if i%2 == 0 {
			w = 2
		}
		pc[id] = map[string][2]int{"x": {w, 2}}
	}
	lo, hi := bootstrapWin("x", ids, pc, 2000, rand.New(rand.NewSource(1)))
	if !(lo < 0.5 && hi > 0.5 && lo > 0.2 && hi < 0.8) {
		t.Fatalf("interval %v–%v", lo, hi)
	}
}
