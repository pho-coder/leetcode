package problems

import (
	"strconv"
)

//countAndSay https://leetcode.com/problems/count-and-say/
func countAndSay(n int) string {
	re := "1"
	for i := 2; i <= n; i++ {
		tmp := string(re[0])
		tmpRe := ""
		for j := 1; j < len(re); j++ {
			if string(tmp[0]) != string(re[j]) {
				tmpRe += string(strconv.Itoa(len(tmp))) + string(tmp[0])
				tmp = string(re[j])
			} else {
				tmp += string(re[j])
			}
		}
		if len(tmp) > 0 {
			re = tmpRe + string(strconv.Itoa(len(tmp))) + string(tmp[0])
		}
	}
	return re
}

func countAndSay1(n int) string {
	re := "1"
	for i := 2; i <= n; i++ {
		tmp := string(re[0])
		strLen := 1
		tmpRe := ""
		for j := 1; j < len(re); j++ {
			jStr := string(re[j])
			if tmp != jStr {
				tmpRe += string(strconv.Itoa(strLen)) + tmp
				tmp = jStr
				strLen = 1
			} else {
				strLen++
			}
		}
		if strLen > 0 {
			re = tmpRe + string(strconv.Itoa(strLen)) + tmp
		}
	}
	return re
}
