package ArrayandString

//spiralOrder https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1168/
func spiralOrder(matrix [][]int) []int {
	re := []int{}
	if len(matrix) == 0 {
		return re
	}
	minI := 0
	minJ := 0
	maxI := len(matrix) - 1
	maxJ := len(matrix[0]) - 1
	i := 0
	j := 0
	direction := "right"
	for {
		re = append(re, matrix[i][j])
		if direction == "right" {
			if j < maxJ {
				j++
			} else if j == maxJ {
				direction = "down"
				if i+1 > maxI {
					break
				} else {
					i++
					minI = i
				}
			}
		} else if direction == "down" {
			if i < maxI {
				i++
			} else if i == maxI {
				direction = "left"
				if j-1 < minJ {
					break
				} else {
					j--
					maxJ = j
				}
			}
		} else if direction == "left" {
			if j > minJ {
				j--
			} else if j == minJ {
				direction = "up"
				if i-1 < minI {
					break
				} else {
					i--
					maxI = i
				}
			}
		} else if direction == "up" {
			if i > minI {
				i--
			} else if i == minI {
				direction = "right"
				if j+1 > maxJ {
					break
				} else {
					j++
					minJ = j
				}
			}
		}
	}
	return re
}
