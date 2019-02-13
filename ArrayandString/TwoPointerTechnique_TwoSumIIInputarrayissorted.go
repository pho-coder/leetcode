package ArrayandString

//twoSum https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1153/
func twoSum(numbers []int, target int) []int {
	index1 := -1
	index2 := -1
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				index1 = i + 1
				index2 = j + 1
				goto Exit
			}
		}
	}
Exit:
	return []int{index1, index2}
}
