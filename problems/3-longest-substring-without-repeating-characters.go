/*

 * 3. Longest Substring Without Repeating Characters
 * Submit Count : 2 times
 * Time used : 1hr 20 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	result := lengthOfLongestSubstring(" ")
	print(fmt.Sprintf("\nresult: %d", result))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	type void struct{}

	// Create unused type
	max := 0
	checklistMapper := make(map[int]map[string]void)

	checkMax := func(cmkKey int) {
		if max < len(checklistMapper[cmkKey]) {
			max = len(checklistMapper[cmkKey])
		}
	}

	for charIdx, char := range strings.Split(s, "") {
		for cmkKey := range checklistMapper {
			checkMax(cmkKey)
			if _, ok := checklistMapper[cmkKey][char]; ok {
				delete(checklistMapper, cmkKey)
				continue
			}
			checklistMapper[cmkKey][char] = void{}
			checkMax(cmkKey)
		}
		checklistMapper[charIdx] = map[string]void{char: {}}
	}

	return max
}
