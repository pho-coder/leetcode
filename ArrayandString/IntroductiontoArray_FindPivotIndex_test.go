package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPivotIndex(t *testing.T) {
	input := []int{1, 7, 3, 6, 5, 6}
	output := 3
	Convey("test pivotIndex", t, func() {
		re := pivotIndex(input)
		So(re, ShouldResemble, output)
	})
}
