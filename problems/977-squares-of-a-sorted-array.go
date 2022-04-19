package main

/*

 * 977. Squares of a Sorted Array
 * Submit Count : 1 time
 * Time used : 6 m

 * Solve by バンクー
 * Code with L#v

 */

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	results := make([]int, 0)

	for _, value := range nums {
		results = append(results, value*value)
	}
	sort.Ints(results)
	return results
}

func main() {
	output := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Printf("%v", output)
}
