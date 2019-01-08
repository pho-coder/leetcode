package problems

//maxSubArray https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	re := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		tmpSum := nums[i]
		if tmpSum > re {
			re = tmpSum
		}
		for j := i + 1; j < len(nums); j++ {
			tmpSum += nums[j]
			if tmpSum > re {
				re = tmpSum
			}
		}
	}
	if nums[len(nums)-1] > re {
		re = nums[len(nums)-1]
	}
	return re
}
