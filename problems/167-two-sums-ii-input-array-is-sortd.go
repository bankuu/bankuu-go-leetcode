/*

 * 167. Two Sum II - Input Array Is Sorted
 * Submit Count : 1 time
 * Time used : 1 m (clip form 1-two-sums)

 * Solve by バンクー
 * Code with L#v

 */

package main

import "fmt"

func main() {
	result := twoSumII([]int{2, 7, 11, 15}, 9)
	fmt.Printf("%v", result)
}

func twoSumII(nums []int, target int) []int {
	for tIdx, thing := range nums {
		for mIdx, matter := range nums[tIdx+1:] {
			if thing+matter == target {
				return []int{tIdx + 1, mIdx + tIdx + 2}
			}
		}
	}
	return []int{}
}
