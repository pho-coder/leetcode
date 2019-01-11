package problems

import "fmt"

//climbStairs https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	//exhaustive method
	re := 0
	// i is count 2
	for i := 0; i*2 <= n; i++ {
		re++
		re += i * (n - 2*i)
		fmt.Println(re)
	}
	return re
}
