package main

func twoSumV2(nums []int, target int) []int {
	var result []int
	m := make(map[int]int)

	for i, k := range nums {
		if value, exits := m[target-k]; exits {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}
