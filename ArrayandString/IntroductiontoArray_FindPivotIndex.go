package ArrayandString

//pivotIndex https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1144/
func pivotIndex(nums []int) int {
	re := -1
	leftSum := map[int]int{}
	for i, v := range nums {
		leftSum[i] = v + leftSum[i-1]
	}
	rightSum := map[int]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		rightSum[i] = nums[i] + rightSum[i+1]
		if rightSum[i] == leftSum[i] {
			re = i
		}
	}
	return re
}
