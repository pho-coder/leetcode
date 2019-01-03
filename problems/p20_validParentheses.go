package problems

// IsValid https://leetcode.com/problems/valid-parentheses/
func IsValid(s string) bool {
	re := true
	noClose := []string{}
Exit:
	for i := range s {
		currStr := s[i : i+1]
		if currStr == "(" || currStr == "{" || currStr == "[" {
			noClose = append(noClose, currStr)
			continue
		}
		switch currStr {
		case ")":
			if len(noClose) == 0 {
				re = false
				break Exit
			} else if noClose[len(noClose)-1] == "(" {
				noClose = noClose[:len(noClose)-1]
				continue
			} else {
				re = false
				break Exit
			}
		case "}":
			if len(noClose) == 0 {
				re = false
				break Exit
			} else if noClose[len(noClose)-1] == "{" {
				noClose = noClose[:len(noClose)-1]
				continue
			} else {
				re = false
				break Exit
			}
		case "]":
			if len(noClose) == 0 {
				re = false
				break Exit
			} else if noClose[len(noClose)-1] == "[" {
				noClose = noClose[:len(noClose)-1]
				continue
			} else {
				re = false
				break Exit
			}
		default:
			re = false
			break Exit
		}
	}
	if len(noClose) > 0 {
		re = false
	}
	return re
}
