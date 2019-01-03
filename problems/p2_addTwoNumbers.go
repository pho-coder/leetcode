package problems

// ListNode list node struct
type ListNode struct {
	Val  int
	Next *ListNode
}

//AddTwoNumbers add two numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nilListNode ListNode
	dummyHead := new(ListNode)
	dummyHead.Val = 0
	dummyHead.Next = &nilListNode
	p := l1
	q := l2
	curr := dummyHead
	carry := 0
	for p != nil || q != nil {
		var x int
		var y int
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = new(ListNode)
		curr.Next.Val = sum % 10
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = new(ListNode)
		curr.Next.Val = carry
	}
	return dummyHead.Next
}
