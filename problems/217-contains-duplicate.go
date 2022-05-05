/*

 * 217. Contains Duplicate
 * Submit Count : 2 time
 * Time used : 39 m

 * Solve by バンクー
 * Code with L#v

 */

package main

func containsDuplicate(nums []int) bool {
	mapper := make(map[int]interface{})
	for _, value := range nums {
		mapper[value] = value
	}
	return len(mapper) == len(nums)
}
