/*

 * 53. Maximum Subarray
 * Submit Count : >10 times
 * Time used :  >3 hours

 * Solve by バンクー
 * Code with L#v

 */

package main

func maxSubArray(nums []int) int {

	maxSums := nums[0]

	sumList := make([]int, 0)

	// for find positive
	for _, value := range nums {

		//set range
		for idx, _ := range sumList {
			sumList[idx] = sumList[idx] + value
		}

		//set range if positive
		if value > 0 {
			sumList = append(sumList, value)
			for _, checkValue := range sumList {
				if checkValue > maxSums {
					maxSums = checkValue
				}
			}

			maxSumList := sumList[0]
			for _, val := range sumList[1:] {
				if maxSumList < val {
					maxSumList = val
				}
			}

			// renew max value
			sumList = []int{maxSumList}
		}

		//set uni position
		if value > maxSums {
			maxSums = value
		}
	}

	return maxSums
}

func main() {
	result := maxSubArray([]int{8, -19, 5, -4, 20})
	print(result)
}
