package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpiralOrder(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	Convey("test spiralOrder", t, func() {
		re := spiralOrder(input)
		So(re, ShouldResemble, output)
	})
}
