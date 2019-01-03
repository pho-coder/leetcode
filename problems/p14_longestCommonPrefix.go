package problems

// LongestCommonPrefix https://leetcode.com/problems/longest-common-prefix/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	re := ""
	for _, astr := range strs {
		if minLen > len(astr) {
			minLen = len(astr)
		}
	}
Exit:
	for i := 0; i < minLen; i++ {
		tmpRe := strs[0][i]
		for _, astr := range strs {
			if astr[i] == tmpRe {
				continue
			} else {
				break Exit
			}
		}
		re = re + string(tmpRe)
	}
	return re
}
