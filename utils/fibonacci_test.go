package utils

import (
	"testing"
	"fmt"
)

func TestSumLargeNumber1(t *testing.T) {
	fmt.Println(string(48))
	fmt.Println(string([]rune{48, 49}))
}

func TestSumLargeNumber(t *testing.T) {
	cases := []struct {
		n1, n2, sum string
	}{
		{"0", "1", "1"},
		{"100", "1", "101"},
		{"10000000", "8766555", "18766555"},
		{"22002056689466296922983322104048463", "13598018856492162040239554477268290", "35600075545958458963222876581316753"},
		{"84885164052257330097714121751630835360966663883732297726369399", "137347080577163115432025771710279131845700275212767467264610201", "222232244629420445529739893461909967206666939096499764990979600"},
	}

	for _, c := range cases {
		got := sumLargeNum(c.n1, c.n2)
		if got != c.sum {
			t.Errorf("%q + %q expected is %q, actual is %q", c.n1, c.n2, c.sum, got)
		}
	}
}
