/*

 * 1. Two Sum
 * Summit Count : 1 time
 * Time used : 13 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import "fmt"

func main() {
	result := twoSum([]int{3, 5, 2, 1}, 6)
	fmt.Printf("%v", result)
}

func twoSum(nums []int, target int) []int {
	for tIdx, thing := range nums {
		for mIdx, matter := range nums[tIdx+1:] {
			if thing+matter == target {
				return []int{tIdx, mIdx + tIdx + 1}
			}
		}
	}
	return []int{}
}
