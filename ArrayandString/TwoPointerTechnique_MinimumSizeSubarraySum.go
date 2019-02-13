package ArrayandString

//minSubArrayLen https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1299/
func minSubArrayLen(s int, nums []int) int {
	re := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum >= s {
			re = 1
			break
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= s {
				if j-i+1 < re {
					re = j - i + 1
				}
				break
			}
		}
	}
	if re == len(nums)+1 {
		re = 0
	}
	return re
}
