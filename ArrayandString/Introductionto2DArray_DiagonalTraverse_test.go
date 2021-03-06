package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindDiagonalOrder(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	Convey("test findDiagonalOrder", t, func() {
		re := findDiagonalOrder(input)
		So(re, ShouldResemble, output)
	})
}
