package problems

//searchInsert https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	re := 0
	found := false
	for i, v := range nums {
		if target <= v {
			re = i
			found = true
			break
		}
	}
	if !found {
		re = len(nums)
	}
	return re
}
