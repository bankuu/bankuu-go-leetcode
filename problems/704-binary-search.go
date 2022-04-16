/*

 * 704. Binary Search
 * Submit Count : 1 time
 * Time used : 21 m

 * Solve by バンクー
 * Code with L#v

 */

package main

func main() {
	result := search([]int{-1, 0, 3, 5, 9, 12}, 5)
	print(result)
}

func search(nums []int, target int) int {
	lPos := 0
	rPos := len(nums) - 1

	checkRes := func(idx int) bool {
		return nums[idx] == target
	}

	for {

		// -- Check result
		if checkRes(lPos) {
			return lPos
		}
		if checkRes(rPos) {
			return rPos
		}
		center := (rPos + lPos) / 2
		if checkRes(center) {
			return center
		}
		// --

		// -- Check not found
		if len(nums[lPos:rPos]) < 3 {
			return -1
		}

		// -- Move Pointer
		if target > nums[center] {
			lPos = center
		} else {
			rPos = center
		}
	}
}
