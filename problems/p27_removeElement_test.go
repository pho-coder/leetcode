package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveElement(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	input2 := 2
	output := 5
	Convey("test RemoveElement", t, func() {
		re := RemoveElement(input, input2)
		So(re, ShouldResemble, output)
	})
}
