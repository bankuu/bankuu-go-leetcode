/*

 * 35. Search Insert Position
 * Submit Count : 1 time
 * Time used : 14 m

 * Solve by バンクー
 * Code with L#v

 */

package main

func searchInsert(nums []int, target int) int {
	lPointer := 0
	rPointer := len(nums)

	for {
		center := (rPointer + lPointer) / 2

		// not found
		if rPointer-lPointer <= 1 {
			if nums[rPointer-1] < target {
				return rPointer
			} else {
				return rPointer - 1
			}
		}

		if nums[center] == target {
			return center
		} else if nums[center] < target {
			lPointer = center
		} else {
			rPointer = center
		}
	}
}

func main() {
	result := searchInsert([]int{1, 3, 5, 6}, 7)
	print(result)
}
