package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNumbers(t *testing.T) {
	var nilListNode ListNode
	a3 := new(ListNode)
	a3.Val = 3
	a3.Next = &nilListNode
	a2 := new(ListNode)
	a2.Val = 2
	a2.Next = a3
	a1 := new(ListNode)
	a1.Val = 1
	a1.Next = a2
	a0 := new(ListNode)
	a0.Val = 0
	a0.Next = a1
	b3 := new(ListNode)
	b3.Val = 4
	b3.Next = &nilListNode
	b2 := new(ListNode)
	b2.Val = 3
	b2.Next = b3
	b1 := new(ListNode)
	b1.Val = 9
	b1.Next = b2
	b0 := new(ListNode)
	b0.Val = 0
	b0.Next = b1
	Convey("test AddTwoNumbers", t, func() {
		re := AddTwoNumbers(a0, b0)
		So(re.Val, ShouldEqual, 0)
		So(re.Next.Val, ShouldEqual, 0)
		So(re.Next.Val, ShouldEqual, 0)
		So(re.Next.Val, ShouldEqual, 0)
	})
}
