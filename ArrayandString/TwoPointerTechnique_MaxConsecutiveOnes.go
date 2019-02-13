package ArrayandString

//findMaxConsecutiveOnes https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1301/
func findMaxConsecutiveOnes(nums []int) int {
	con := 0
	maxCon := 0
	for _, one := range nums {
		if one == 1 {
			con += 1
			if con > maxCon {
				maxCon = con
			}
		} else {
			con = 0
		}
	}
	return maxCon
}
