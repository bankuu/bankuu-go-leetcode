package main

/*

 * 1480. Running Sum of 1d Array
 * Submit Count : 1 time
 * Time used : 4 minutes

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

func runningSum(nums []int) []int {
	for idx := 1; idx < len(nums); idx++ {
		nums[idx] = nums[idx-1] + nums[idx]
	}
	return nums
}

func main() {
	fmt.Printf("%v", runningSum([]int{1, 2, 3, 4}))
}
