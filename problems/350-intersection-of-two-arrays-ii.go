package main

/*

 * 88. Merge Sorted Array
 * Submit Count : 1 time
 * Time used : 20 m

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {

	result := make([]int, 0)
	var thing *[]int
	var matter *[]int

	if len(nums1) < len(nums2) {
		thing = &nums1
		matter = &nums2
	} else {
		thing = &nums2
		matter = &nums1
	}

	for _, dataThing := range *thing {
		for idxMatter, dataMatter := range *matter {
			if dataThing == dataMatter {
				result = append(result, dataMatter)
				*matter = append((*matter)[:idxMatter], (*matter)[idxMatter+1:]...)
				break
			}
		}
	}
	return result
}

func main() {
	num1 := []int{4, 9, 5}
	num2 := []int{9, 4, 9, 8, 4}

	result := intersect(num1, num2)
	fmt.Printf("%v", result)
}
