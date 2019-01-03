package problems

// IsPalindrome https://leetcode.com/problems/palindrome-number/
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	re := 0
	y := x
	for y != 0 {
		tmp := y % 10
		y = y / 10
		re = re*10 + tmp
	}
	if re == x {
		return true
	}
	return false
}
