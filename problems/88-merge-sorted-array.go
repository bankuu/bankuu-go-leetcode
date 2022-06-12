/*

 * 88. Merge Sorted Array
 * Submit Count : 1 times
 * Time used :  10 minutes

 * Solve by バンクー
 * Code with L#v

 */

package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	// to list
	list := nums1[:m]
	list = append(list, nums2[:n]...)

	// to sort order
	sort.Ints(list)

	// for num1
	for idx, _ := range nums1 {
		nums1[idx] = list[idx]
	}

}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	num2 := []int{2, 5, 6}
	n := 3
	merge(num1, n, num2, m)
	fmt.Printf("%v", num1)
}
