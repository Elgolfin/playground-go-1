package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var wordCount = make(map[string]int)
	var words = strings.Fields(s)
	for _, word := range words {
		_, ok := wordCount[word]
		if ok {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
