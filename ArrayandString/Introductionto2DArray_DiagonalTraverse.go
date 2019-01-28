package ArrayandString

// findDiagonalOrder https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1167/
func findDiagonalOrder(matrix [][]int) []int {
	re := []int{}
	if len(matrix) == 0 {
		return re
	}
	m := len(matrix)
	n := len(matrix[0])
	maxLen := m
	if n < m {
		maxLen = n
	}
	maxLenstep := m + n - 1 - maxLen + 1
	i := 0
	j := 0
	direction := true
	lenNow := 1
	for step := 1; step <= m+n-1; step++ {
		stepOne := 0
		for {
			re = append(re, matrix[i][j])
			stepOne++
			if stepOne >= lenNow {
				break
			}
			if direction {
				i--
				j++
			} else {
				i++
				j--
			}
		}
		if direction && j < n-1 {
			j++
		} else if direction && j == n-1 {
			i++
		} else if !direction && i < m-1 {
			i++
		} else if !direction && i == m-1 {
			j++
		}
		if step < maxLenstep && lenNow < maxLen {
			lenNow++
		} else if step >= maxLenstep {
			lenNow--
		}
		direction = !direction
	}
	return re
}
