package problems

//RemoveDuplicates https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	re := 0
	if len(nums) == 0 {
		return re
	}
	o := nums[0]
	re = 1
	for _, one := range nums {
		if one > o {
			o = one
			nums[re] = o
			re++
		}
	}
	return re
}
