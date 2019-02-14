package ArrayandString

//getRow https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1171/
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	} else {
		lastRow := getRow(rowIndex - 1)
		re := []int{1}
		for i := 1; i < rowIndex; i++ {
			re = append(re, lastRow[i-1]+lastRow[i])
		}
		re = append(re, 1)
		return re
	}
}
