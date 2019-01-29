package ArrayandString

//generate https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1170/
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	re := []int{}
	lastRe := generate(numRows - 1)
	lastRow := lastRe[len(lastRe)-1]
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			re = append(re, 1)
		} else {
			re = append(re, lastRow[i-1]+lastRow[i])
		}
	}
	return append(lastRe, re)
}
