package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxSubArray(t *testing.T) {
	input1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := 6
	Convey("test maxSubArray", t, func() {
		re := maxSubArray(input1)
		So(re, ShouldResemble, output)
	})
}
