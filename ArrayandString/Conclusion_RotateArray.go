package ArrayandString

//rotate https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1182/
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}
