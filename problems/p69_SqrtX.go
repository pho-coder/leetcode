package problems

//mySqrt https://leetcode.com/problems/sqrtx/
func mySqrt(x int) int {
	found := false
	start := 0
	middle := 0
	end := x
	re := 0
	for !found {
		middle = (start + end) / 2
		if middle*middle <= x && (middle+1)*(middle+1) >= x {
			found = true
			if (middle+1)*(middle+1) == x {
				re = middle + 1
			} else {
				re = middle
			}
		} else if (middle+1)*(middle+1) < x {
			start = middle
		} else if middle*middle > x {
			end = middle
		}
	}
	return re
}
