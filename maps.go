package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	a := strings.Split(s, " ")
	wordMap := make(map[string]int)
	for _, value := range a {
		wordMap[value]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
