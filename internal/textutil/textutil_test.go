package textutil

import "testing"

// 代码里有中文时，遮盖后的文本必须和原文的 byte 数相同，后面命中的位置才对得上。
func TestMaskCodeKeepsByteOffsets(t *testing.T) {
	s := "先跑 `检查一下` 这个命令。\n值得注意的是，结果。"
	m := MaskCode(s)
	if len(m) != len(s) {
		t.Fatalf("长度变了：原文 %d byte，遮盖后 %d byte", len(s), len(m))
	}
	want := "值得注意的是"
	i := len("先跑 `检查一下` 这个命令。\n")
	if m[i:i+len(want)] != want {
		t.Fatalf("遮盖后偏移错位：%q", m[i:i+len(want)])
	}
	if LineOf(m, i) != 2 {
		t.Fatalf("行号应为 2，得到 %d", LineOf(m, i))
	}
}
