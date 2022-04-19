/*

 * 189. Rotate Array
 * Submit Count : 2 time
 * Time used : 39 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import "fmt"

func rotate(nums []int, k int) {
	oldNums := make([]int, len(nums))
	copy(oldNums, nums)
	k = k % len(nums)
	for idx, _ := range oldNums {
		if k+idx <= len(nums)-1 {
			nums[k+idx] = oldNums[idx]
		} else {
			nums[(k+idx)-(len(nums))] = oldNums[idx]
		}
	}
}

func main() {
	nums := []int{-1}
	rotate(nums, 2)
	fmt.Printf("%v", nums)
}
