package strutil

import "testing"

func TestReverse(t *testing.T) {
	cases := map[string]string{
		"abc":   "cba",
		"你好":   "好你",
		"":      "",
		"a b c": "c b a",
	}
	for in, want := range cases {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%q) = %q，期望 %q", in, got, want)
		}
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank("") || !IsBlank("  \t\n") {
		t.Error("空白字符串应判定为 blank")
	}
	if IsBlank("x") {
		t.Error("非空字符串不应判定为 blank")
	}
}
