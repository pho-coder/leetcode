package main

import (
	"fmt"

	"github.com/pho-coder/leetcode/problems"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(problems.TwoSum(nums, target))
}
