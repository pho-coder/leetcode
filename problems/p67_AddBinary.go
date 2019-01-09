package problems

import "strconv"

//addBinary https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	re := ""
	tmp := 0
	for i >= 0 || j >= 0 || tmp > 0 {
		x := 0
		y := 0
		if i >= 0 {
			x, _ = strconv.Atoi(string(a[i]))
			i--
		}
		if j >= 0 {
			y, _ = strconv.Atoi(string(b[j]))
			j--
		}
		s := tmp + x + y
		if s == 3 {
			re = "1" + re
			tmp = 1
		} else if s == 2 {
			re = "0" + re
			tmp = 1
		} else if s == 1 {
			re = "1" + re
			tmp = 0
		} else if s == 0 {
			re = "0" + re
			tmp = 0
		}
	}
	return re
}
