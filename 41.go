package main

import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	word_count := make(map[string]int)

	var fileds = strings.Fields(s)
	for _, word := range fileds {
		v, is_word := word_count[word]
		if !is_word {
			word_count[word] = 1
		} else {
			word_count[word] = v + 1
		}
	}
	return word_count
}

func main() {
	wc.Test(WordCount)
}
