package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlusOne(t *testing.T) {
	input1 := []int{9, 9, 9}
	output := []int{1, 0, 0, 0}
	Convey("test plusOne", t, func() {
		re := plusOne(input1)
		So(re, ShouldResemble, output)
	})
}
