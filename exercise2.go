package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.

func WordCount(s string) map[string]int {
	stringArr := strings.Split(s, " ")
	m := make(map[string]int)
	for _, value := range stringArr {
		m[value]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}