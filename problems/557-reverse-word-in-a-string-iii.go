package main

/*

 * 557. Reverse Words in a String III
 * Submit Count : 1 time
 * Time used : 12 m

 * Solve by バンクー
 * Code with L#v

 */

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	reverseString := func(s []byte) {
		pointer := 0
		for pointer < (len(s) / 2) {
			s[pointer], s[len(s)-pointer-1] = s[len(s)-pointer-1], s[pointer]
			pointer++
		}
	}

	words := strings.Split(s, " ")
	for idx, word := range words {
		chars := []byte(word)
		reverseString(chars)
		words[idx] = string(chars)
	}

	return strings.Join(words, " ")
}

func main() {
	input := "Let's take LeetCode contest"
	output := reverseWords(input)
	fmt.Printf("%v", output)
}
