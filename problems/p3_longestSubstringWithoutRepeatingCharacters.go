package problems

// LengthOfLongestSubstring1 LengthOfLongestSubstring
func LengthOfLongestSubstring1(s string) int {
	tmpHashMap := map[rune]int{}
	minIndex := 0
	maxIndex := 0
	maxLen := 0
	for i, c := range s {
		if val, ok := tmpHashMap[c]; ok && val >= minIndex {
			minIndex = val + 1
		}
		tmpHashMap[c] = i
		maxIndex = i
		if (maxIndex - minIndex + 1) > maxLen {
			maxLen = maxIndex - minIndex + 1
		}
	}
	return maxLen
}
