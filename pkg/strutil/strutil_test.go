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
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank("") || !IsBlank("  \t\n") {
		t.Error("whitespace-only strings should be blank")
	}
	if IsBlank("x") {
		t.Error("non-empty strings should not be blank")
	}
}
