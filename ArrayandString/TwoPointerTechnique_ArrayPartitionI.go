package ArrayandString

import "sort"

//arrayPairSum https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1154/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	re := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			re += nums[i]
		}
	}
	return re
}
