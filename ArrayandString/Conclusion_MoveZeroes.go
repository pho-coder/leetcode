package ArrayandString

//moveZeroes https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1174/
func moveZeroes(nums []int) {
	zerosNums := 0
	for startIndex := 0; startIndex < len(nums)-zerosNums; {
		if nums[startIndex] == 0 {
			for j := startIndex; j < len(nums)-1-zerosNums; j++ {
				nums[j] = nums[j+1]
			}
			zerosNums++
		} else {
			startIndex++
		}
	}
	for i := len(nums) - zerosNums; i < len(nums); i++ {
		nums[i] = 0
	}
}
