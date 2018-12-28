package problems

//RomanToInt https://leetcode.com/problems/roman-to-integer/
func RomanToInt(s string) int {
	sLen := len(s)
	index := 0
	re := 0
	for i := range s {
		if index > i {
			continue
		}
		if i < sLen-1 {
			switch tmpStr := s[i : i+2]; tmpStr {
			case "IV":
				re = re + 4
				index = index + 2
			case "IX":
				re = re + 9
				index = index + 2
			case "XL":
				re = re + 40
				index = index + 2
			case "XC":
				re = re + 90
				index = index + 2
			case "CD":
				re = re + 400
				index = index + 2
			case "CM":
				re = re + 900
				index = index + 2
			}

		}
		if index > i {
			continue
		}
		switch tmpStr := s[i : i+1]; tmpStr {
		case "I":
			re = re + 1
			index = index + 1
		case "V":
			re = re + 5
			index = index + 1
		case "X":
			re = re + 10
			index = index + 1
		case "L":
			re = re + 50
			index = index + 1
		case "C":
			re = re + 100
			index = index + 1
		case "D":
			re = re + 500
			index = index + 1
		case "M":
			re = re + 1000
			index = index + 1
		}
	}
	return re
}
