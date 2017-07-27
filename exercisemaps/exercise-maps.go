package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount this func is an implementation of
// "A Tour of Go" exercise maps
func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	fieldCount := map[string]int{}
	for _, val := range fields {
		if v, ok := fieldCount[val]; ok {
			fieldCount[val] = v + 1
		} else {
			fieldCount[val] = 1
		}
	}
	return fieldCount
}

func main() {
	wc.Test(WordCount)
}
