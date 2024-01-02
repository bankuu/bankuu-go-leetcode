package solutions

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsLen := len(nums1) + len(nums2)
	numPos := float64(numsLen) / 2

	var sortNum []int
	for {
		var num1 *int
		var num2 *int
		if len(nums1) > 0 {
			num1 = &nums1[0]
		}
		if len(nums2) > 0 {
			num2 = &nums2[0]
		}

		if num1 == nil && num2 == nil || len(sortNum) > int(math.Trunc(numPos)) {
			break
		} else if (num1 != nil && num2 != nil && *num1 < *num2) || (num1 != nil && num2 == nil) {
			sortNum = append(sortNum, *num1)
			nums1 = nums1[1:]
		} else {
			sortNum = append(sortNum, *num2)
			nums2 = nums2[1:]
		}
	}

	if math.Trunc(numPos) == numPos {
		return float64(sortNum[len(sortNum)-2]+sortNum[len(sortNum)-1]) / 2
	} else {
		return float64(sortNum[len(sortNum)-1])
	}
}
