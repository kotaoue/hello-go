package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	find := make(map[string]int)
	for _, word := range strings.Fields(s) {
		find[word]++
	}
	return find
}

func main() {
	wc.Test(WordCount)
}
