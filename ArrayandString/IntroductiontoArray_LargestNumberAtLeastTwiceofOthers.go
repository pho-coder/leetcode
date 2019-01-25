package ArrayandString

//dominantIndex https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1147/
func dominantIndex(nums []int) int {
	re := -1
	firstLargest := 0
	firstIndex := 0
	secondLarget := 0
	for i, v := range nums {
		if v >= firstLargest {
			secondLarget = firstLargest
			firstLargest = v
			firstIndex = i
		} else if v >= secondLarget {
			secondLarget = v
		}
	}
	if firstLargest >= secondLarget*2 {
		re = firstIndex
	}
	return re
}
