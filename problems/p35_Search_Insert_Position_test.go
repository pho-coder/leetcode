package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchInsert(t *testing.T) {
	input1 := []int{1, 3, 5, 6}
	input2 := 5
	output := 2
	Convey("test searchInsert", t, func() {
		re := searchInsert(input1, input2)
		So(re, ShouldResemble, output)
	})
}
