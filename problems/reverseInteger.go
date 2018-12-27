package problems

import "math"

// Reverse https://leetcode.com/problems/reverse-integer/
func Reverse(x int) int {
	re := 0
	if x > math.MaxInt32 || x < math.MinInt32 {
		return re
	}
	for x != 0 {
		tmp := x % 10
		x = x / 10
		re = re*10 + tmp
	}
	if re > math.MaxInt32 || re < math.MinInt32 {
		return 0
	}
	return re
}
