/*

 * 13. Roman to Integer
 * Submit Count : 1 time
 * Time used : 20 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	romanToInt("IV")
}

func romanToInt(input string) int {
	result := 0
	table := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
		"IV": -2, "IX": -2, "XL": -20, "XC": -20, "CD": -200, "CM": -200,
	}
	for index, char := range strings.Split(input, "") {
		result = result + table[char]
		if index > 0 {
			result = result + table[fmt.Sprintf("%c%c", input[index-1], input[index])]
		}
	}
	return result
}
