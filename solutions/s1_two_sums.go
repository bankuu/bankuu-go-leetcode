package solutions

func twoSum(nums []int, target int) []int {
	for matter := range nums {
		for then := range nums[matter+1:] {
			next := then + matter + 1
			if nums[matter]+nums[next] == target {
				return []int{matter, next}
			}
		}
	}
	return []int{}
}
