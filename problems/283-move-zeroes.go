/*

 * 283. Move Zeroes
 * Submit Count : 1 times
 * Time used : 16 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	lCursor := 0
	rCursor := 1
	for {
		if rCursor >= len(nums) {
			return
		}
		if nums[lCursor] == 0 && nums[rCursor] == 0 {
			rCursor++
			continue
		}
		if nums[lCursor] != 0 {
			lCursor++
			rCursor++
			continue
		}
		if nums[rCursor] != 0 {
			nums[lCursor] = nums[rCursor]
			nums[rCursor] = 0
		}
	}
}

func main() {
	input := []int{1, 1, 2, 0, 1, 1}
	moveZeroes(input)
	fmt.Printf("%v", input)
}
