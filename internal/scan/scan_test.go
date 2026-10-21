package scan

import "testing"

func TestHalfWidth(t *testing.T) {
	r := Scan("我们完成了压测,结果不错.下一步:上线;然后观察。")
	if len(r.HalfWidth) != 4 {
		t.Fatalf("want 4 half-width hits, got %d: %+v", len(r.HalfWidth), r.HalfWidth)
	}
}

func TestHalfWidthExceptions(t *testing.T) {
	text := "版本 Go 1.22，共 1,000 条，21:00 开始。见 `a.b;c` 和 https://x.com/a,b 以及 config.yaml。He said: hello, world."
	r := Scan(text)
	if len(r.HalfWidth) != 0 {
		t.Fatalf("want no hits, got %+v", r.HalfWidth)
	}
	r = Scan("```\nfoo(a, b); bar.baz()\n```\n正文。")
	if len(r.HalfWidth) != 0 {
		t.Fatalf("code block should be skipped, got %+v", r.HalfWidth)
	}
}

func TestEllipsisOnce(t *testing.T) {
	r := Scan("还有很多...")
	if len(r.HalfWidth) != 1 {
		t.Fatalf("want 1, got %+v", r.HalfWidth)
	}
}

func TestMeta(t *testing.T) {
	r := Scan("这个现象值得深思。不难看出，问题在流程。")
	if len(r.Meta) != 2 {
		t.Fatalf("want 2 meta hits, got %+v", r.Meta)
	}
}

func TestNotButIs(t *testing.T) {
	text := "问题不是预算，而是优先级。真正决定成败的是执行。看似简单，实则复杂。"
	r := Scan(text)
	if len(r.NotButIs) != 3 || !r.NotButOver() {
		t.Fatalf("want 3 hits over limit, got %d", len(r.NotButIs))
	}
	r = Scan("不是网络的问题，是证书过期了。")
	if len(r.NotButIs) != 1 || r.NotButOver() {
		t.Fatalf("one occurrence is allowed, got %+v", r.NotButIs)
	}
}

func TestNumbering(t *testing.T) {
	text := "# 年度复盘——我的思考\n\n## 一、背景\n\n正文。\n\n## 增长三板斧\n\n第一刀砍需求。\n"
	r := Scan(text)
	if len(r.TitleDash) != 1 {
		t.Fatalf("want title dash, got %+v", r.TitleDash)
	}
	kinds := map[string]int{}
	for _, h := range r.Numbering {
		kinds[h.Kind]++
	}
	if kinds["机械编号"] != 1 || kinds["量词标题"] != 1 || kinds["量词标签"] != 1 {
		t.Fatalf("unexpected numbering hits: %+v", r.Numbering)
	}
}

func TestCleanText(t *testing.T) {
	text := "# 缓存过期时间设成一样会打爆数据库\n\n上周三晚上，商品详情页的缓存同时过期，数据库连接数在一分钟内打满。我们把过期时间加了随机偏移，之后没再出现。\n"
	r := Scan(text)
	if !r.Clean() {
		t.Fatalf("want clean, got %+v", r)
	}
}

func TestPlainOrdinalTitle(t *testing.T) {
	r := Scan("一、项目背景\n\n正文。\n\n1. 打开设置\n2. 点保存\n")
	if len(r.Numbering) != 1 {
		t.Fatalf("want 1 (plain ordinal title only), got %+v", r.Numbering)
	}
}

func TestSkipQuoted(t *testing.T) {
	text := "元话语比如“值得深思”要删。这件事值得关注。"
	if n := len(ScanWith(text, Options{SkipQuoted: true}).Meta); n != 1 {
		t.Fatalf("want 1 unquoted meta hit, got %d", n)
	}
}
