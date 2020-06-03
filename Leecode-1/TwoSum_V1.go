package main

func twoSumV1(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	for i , v := range nums {
		for k := i + 1; k < len(nums); k ++ {
			if target - v == nums[k] {
				return []int{i, k}
			}
		}
	}

	return []int{}
}