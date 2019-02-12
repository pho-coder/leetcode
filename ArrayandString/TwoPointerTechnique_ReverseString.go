package ArrayandString

//reverseString https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1183/
func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		tmp := s[i]
		s[i] = s[l-1-i]
		s[l-1-i] = tmp
	}
}
