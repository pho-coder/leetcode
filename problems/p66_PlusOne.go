package problems

//plusOne https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	re := digits
	tmp := 1
	for i := len(digits) - 1; i >= 0; i-- {
		t := digits[i] + tmp
		if t >= 10 {
			re[i] = t - 10
			tmp = 1
		} else {
			re[i] = t
			tmp = 0
		}
	}
	if tmp > 0 {
		re = append([]int{1}, re...)
	}
	return re
}
