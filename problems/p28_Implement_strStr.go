package problems

// strStr https://leetcode.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	re := -1
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen == 0 {
		return 0
	}
	for i, v := range haystack {
		if re == -1 && string(needle[0]) == string(v) {
			if haystackLen-i < needleLen {
				return re
			} else if haystack[i:i+needleLen] == needle {
				return i
			}
		}
	}
	return re
}
