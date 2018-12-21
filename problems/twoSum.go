package problems

// TwoSum1 https://leetcode.com/problems/two-sum/
func TwoSum1(nums []int, target int) []int {
	re := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				re = []int{i, j}
				return re
			}
		}
	}
	return re
}

// TwoSum2 https://leetcode.com/problems/two-sum/
func TwoSum2(nums []int, target int) []int {
	tmpMap := map[int]int{}
	re := []int{}
	for i := 0; i < len(nums); i++ {
		tmpMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if val, ok := tmpMap[tmp]; ok {
			if val != i {
				re = []int{i, val}
				return re
			}
		}
	}
	return re
}
