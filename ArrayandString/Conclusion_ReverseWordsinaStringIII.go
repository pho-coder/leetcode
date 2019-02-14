package ArrayandString

import "strings"

//reverseWords https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1165/
func reverseWords(s string) string {
	re := ""
	ss := strings.Split(s, " ")
	for _, one := range ss {
		tmp := []byte{}
		for j := len(one) - 1; j >= 0; j-- {
			tmp = append(tmp, one[j])
		}
		tmpStr := string(tmp)
		if re == "" {
			re = tmpStr
		} else {
			re = re + " " + tmpStr
		}
	}
	return re
}
