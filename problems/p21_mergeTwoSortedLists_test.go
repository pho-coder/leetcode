package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoLists(t *testing.T) {
	var nilListNode ListNode
	// 1 2 3
	a3 := new(ListNode)
	a3.Val = 3
	a3.Next = &nilListNode
	a2 := new(ListNode)
	a2.Val = 2
	a2.Next = a3
	a1 := new(ListNode)
	a1.Val = 1
	a1.Next = a2
	// 2 3 4
	b3 := new(ListNode)
	b3.Val = 4
	b3.Next = &nilListNode
	b2 := new(ListNode)
	b2.Val = 3
	b2.Next = b3
	b1 := new(ListNode)
	b1.Val = 2
	b1.Next = b2
	// 1 2 2 3 3 4
	Convey("test MergeTwoLists", t, func() {
		re := MergeTwoLists(a1, b1)
		So(re.Val, ShouldEqual, 1)
		So(re.Next.Val, ShouldEqual, 2)
		So(re.Next.Next.Val, ShouldEqual, 2)
		So(re.Next.Next.Next.Val, ShouldEqual, 3)
	})
}
