package problems

//lengthOfLastWord https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	lastLength := 0
	tmpLength := 0
	for _, val := range s {
		if " " == string(val) {
			if tmpLength > 0 {
				lastLength = tmpLength
			}
			tmpLength = 0
		} else {
			tmpLength++
		}
	}
	if tmpLength > 0 {
		lastLength = tmpLength
	}
	return lastLength
}
