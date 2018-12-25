package main

import (
	"fmt"

	"github.com/pho-coder/leetcode/problems"
)

func main() {
	var nilListNode problems.ListNode
	a3 := new(problems.ListNode)
	a3.Val = 3
	a3.Next = &nilListNode
	a2 := new(problems.ListNode)
	a2.Val = 2
	a2.Next = a3
	a1 := new(problems.ListNode)
	a1.Val = 1
	a1.Next = a2
	a0 := new(problems.ListNode)
	a0.Val = 0
	a0.Next = a1
	b3 := new(problems.ListNode)
	b3.Val = 4
	b3.Next = &nilListNode
	b2 := new(problems.ListNode)
	b2.Val = 3
	b2.Next = b3
	b1 := new(problems.ListNode)
	b1.Val = 9
	b1.Next = b2
	b0 := new(problems.ListNode)
	b0.Val = 0
	b0.Next = b1
	re := problems.AddTwoNumbers(a0, b0)
	fmt.Println(re.Val)
	fmt.Println(re.Next.Val)
	fmt.Println(re.Next.Next.Val)
	fmt.Println(re.Next.Next.Next.Val)
}
