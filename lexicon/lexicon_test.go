package lexicon

import "testing"

func TestBuiltinCompiles(t *testing.T) {
	l, err := Default()
	if err != nil {
		t.Fatal(err)
	}
	if len(l.Rules) < 100 {
		t.Errorf("规则太少：%d", len(l.Rules))
	}
	for _, r := range l.Rules {
		if r.Hint == "" {
			t.Errorf("规则 %s 缺少 hint", r.ID)
		}
		if r.Regexp().MatchString("") {
			t.Errorf("规则 %s 会匹配空串", r.ID)
		}
	}
}
