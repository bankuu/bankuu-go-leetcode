package main

/*

 * 224. Reverse String
 * Submit Count : 2 times
 * Time used : 25 m

 * Solve by バンクー
 * Code with L#v

 */

import (
	"fmt"
)

func reverseString(s []byte) {
	pointer := 0
	for pointer < (len(s) / 2) {
		s[pointer], s[len(s)-pointer-1] = s[len(s)-pointer-1], s[pointer]
		pointer++
	}
}

func main() {
	input := []byte("hello")
	reverseString(input)
	fmt.Printf("%v", string(input))
}
