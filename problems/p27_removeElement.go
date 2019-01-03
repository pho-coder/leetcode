package problems

// RemoveElement https://leetcode.com/problems/remove-element/
func RemoveElement(nums []int, val int) int {
	re := 0
	for _, one := range nums {
		if one != val {
			nums[re] = one
			re++
		}
	}
	return re
}
