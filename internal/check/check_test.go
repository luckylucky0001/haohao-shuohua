package check

import (
	"testing"

	"haohao-shuohua/lexicon"
)

func TestMissingAndAdded(t *testing.T) {
	orig := "峰值 QPS 达到 1.2 万，较上周提升了 40%，计划 11 月 8 日前完成，见 `SESSION_TTL`。"
	rw := "峰值 QPS 1.2万，比上周高 50%，11 月 8 日前完成。"
	res := Compare(lexicon.MustDefault(), orig, rw)
	missing := map[string]bool{}
	for _, m := range res.Missing {
		missing[m.Text] = true
	}
	if !missing["40%"] || !missing["`SESSION_TTL`"] {
		t.Errorf("应发现 40%% 和代码片段丢失，得到 %+v", res.Missing)
	}
	if missing["1.2 万"] {
		t.Errorf("“1.2 万”与“1.2万”只是空格不同，不算丢失")
	}
	added := false
	for _, a := range res.Added {
		if a.Text == "50%" {
			added = true
		}
	}
	if !added {
		t.Errorf("应发现新增的 50%%，得到 %+v", res.Added)
	}
}

func TestResidualAndIntroduced(t *testing.T) {
	orig := "值得注意的是，我们有效赋能了一线同学。"
	rw := "说白了，我们赋能了一线同学。"
	res := Compare(lexicon.MustDefault(), orig, rw)
	if len(res.Residual) == 0 {
		t.Errorf("“赋能”应算残留")
	}
	if res.EditRatio <= 0 {
		t.Errorf("改动率应大于 0")
	}
}

func TestUnchanged(t *testing.T) {
	text := "老王，明天下午的评审我去不了。PPT 在群文件，第 5 页还差华东的数据。"
	res := Compare(lexicon.MustDefault(), text, text)
	if len(res.Missing) != 0 || len(res.Added) != 0 || res.EditRatio != 0 {
		t.Errorf("原样返回不应有差异：%+v", res)
	}
}

func TestChineseNumeralsMatchDigits(t *testing.T) {
	orig := "团队从 40 人壮大到 95 人，营收同比增长 68%，接起了 1.3 万通电话，连续加班七天。"
	rw := "团队从四十个人变成九十五人，营收增长了百分之六十八，接了一万三千通电话，连着加了 7 天班。"
	res := Compare(lexicon.MustDefault(), orig, rw)
	for _, m := range res.Missing {
		if m.Text == "68%" || m.Text == "1.3 万" || m.Text == "95 人" {
			t.Errorf("中文数字写法不应算丢失：%+v", m)
		}
	}
}

func TestCnToInt(t *testing.T) {
	cases := map[string]int64{"九十五": 95, "四十": 40, "十一": 11, "一万三千": 13000, "两": 2, "三千": 3000, "一百零五": 105}
	for in, want := range cases {
		if got := cnToInt(in); got != want {
			t.Errorf("cnToInt(%q)=%d，应为 %d", in, got, want)
		}
	}
}
