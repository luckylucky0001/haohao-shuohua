package lint

import (
	"strings"
	"testing"

	"haohao-shuohua/lexicon"
)

func hitsOf(t *testing.T, text string) []Hit {
	t.Helper()
	return Lint(lexicon.MustDefault(), text).Hits
}

func hasRule(hits []Hit, id string) bool {
	for _, h := range hits {
		if h.RuleID == id {
			return true
		}
	}
	return false
}

func TestCatchesTypicalPatterns(t *testing.T) {
	cases := map[string]string{
		"值得注意的是，缓存已经失效。":          "op-zhide",
		"在当今快速发展的云原生时代，可观测性很重要。":  "op-dangjin",
		"这套机制有效赋能一线值班同学。":         "jg-funeng",
		"答案其实很简单：方向错了。":           "tp-daan",
		"希望以上建议对你有所帮助！":           "cl-xiwang",
		"希望这封邮件找到您一切安好！":          "tr-email",
		"先把线上稳稳接住。":               "en-wenwen",
		"这家店真的绝绝子！":               "ws-juejuezi",
		"分布式链路追踪扮演着至关重要的角色。":      "lf-banyan",
		"我们已第一时间对相关流程进行了全面排查和优化。": "tr-dui-jinxing",
		"让我们携手共进，共同打造更加智能的办公环境！":  "cl-xieshou",
		"本研究具有重要的理论意义和实践价值。":      "lf-yiyi",
		"不是你不够努力，而是你努力的方向错了。":     "tp-bushi-ershi",
		"大理，一座被时光温柔以待的城市。":        "lt-wenrou",
		"您的满意是我们最大的追求。":           "sv-manyi",
	}
	for text, id := range cases {
		if hits := hitsOf(t, text); !hasRule(hits, id) {
			t.Errorf("%q 应命中 %s，实际 %v", text, id, hits)
		}
	}
}

func TestNaturalTextIsQuiet(t *testing.T) {
	natural := []string{
		"老王，明天下午的方案评审我去不了，孩子学校临时要开家长会。PPT 我已经传到群文件了。",
		"周日带我妈去复查，医生说指标比上次好，药可以减半。",
		"各部门：根据集团年度安全检查安排，定于 11 月 3 日至 11 月 7 日开展消防安全专项检查。特此通知。",
	}
	for _, text := range natural {
		rep := Lint(lexicon.MustDefault(), text)
		for _, h := range rep.Hits {
			if h.W >= 2 {
				t.Errorf("自然文本不应有权重 ≥2 的命中：%q 命中 %s（%s）", text, h.RuleID, h.Match)
			}
		}
	}
}

func TestCodeIsMasked(t *testing.T) {
	text := "运行 `赋能 --all` 之后看日志。\n```\n// 值得注意的是：这里是代码注释\n```\n"
	if hits := hitsOf(t, text); len(hits) != 0 {
		t.Errorf("代码里的内容不应命中：%v", hits)
	}
}

func TestStats(t *testing.T) {
	st := Lint(lexicon.MustDefault(), "## 标题\n- ✅ 第一\n- ✅ 第二\n**加粗**很好！高效便捷、安全可靠、灵活多样。是每一位伙伴的付出，是每一个深夜的坚守，是每一次困难面前的不放弃。").Stats
	if st.Headings != 1 || st.ListLines != 2 || st.Emoji != 2 || st.Bold != 1 || st.Exclaims != 1 {
		t.Errorf("排版统计不对：%+v", st)
	}
	if st.FourCharRuns != 1 {
		t.Errorf("应识别一处四字格连用，得到 %d", st.FourCharRuns)
	}
	if len(st.Parallels) != 1 || !strings.HasPrefix(st.Parallels[0], "是每一位") {
		t.Errorf("应识别排比，得到 %v", st.Parallels)
	}
}

func TestNoFalseFourCharRun(t *testing.T) {
	st := Lint(lexicon.MustDefault(), "我说外卖半小时就到，她不理我，拎着一把葱、两斤排骨回来。").Stats
	if st.FourCharRuns != 0 {
		t.Errorf("长分句不应算四字格连用，得到 %d", st.FourCharRuns)
	}
}

func TestToolTracesInLinks(t *testing.T) {
	text := "详见 https://example.com/post?utm_source=chatgpt.com 和 turn0search3。作为一个 AI 语言模型，我无法确认。"
	hits := hitsOf(t, text)
	for _, id := range []string{"tc-utm", "tc-turn", "tc-selfref"} {
		if !hasRule(hits, id) {
			t.Errorf("应命中 %s，实际 %v", id, hits)
		}
	}
}
