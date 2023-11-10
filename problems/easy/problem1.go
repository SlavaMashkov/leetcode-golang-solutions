package easy

// 1. Two Sum
// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	diffMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if index, contains := diffMap[diff]; contains {
			return []int{index, i}
		}

		diffMap[num] = i
	}

	return []int{0, 0}
}
