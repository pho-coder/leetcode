package problems

//MergeTwoLists https://leetcode.com/problems/merge-two-sorted-lists/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Val = 0
	p := l1
	q := l2
	curr := dummyHead
	for p != nil || q != nil {
		var x int
		var y int
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		curr.Next = new(ListNode)
		if p != nil && q != nil {
			if x <= y {
				curr.Next.Val = x
				p = p.Next
			} else {
				curr.Next.Val = y
				q = q.Next
			}
		} else if p != nil {
			curr.Next.Val = x
			p = p.Next
		} else if q != nil {
			curr.Next.Val = y
			q = q.Next
		}
		curr = curr.Next
	}
	return dummyHead.Next
}
