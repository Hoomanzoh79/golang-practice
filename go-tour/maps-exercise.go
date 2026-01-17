package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		counts[w]++
	}
	return counts
}
