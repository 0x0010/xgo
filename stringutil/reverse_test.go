package stringutil

import (
	"testing"
	"fmt"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestConversion(t *testing.T) {
	//rune type
	codePoint := int32(22909)
	// string value
	str := "中"
	// code point of "中" is 22909
	if codePoint != []rune(str)[0] {
		fmt.Errorf("%v rune value is not %d", str, codePoint)
	}
}
