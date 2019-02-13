package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	output := []int{5, 6, 7, 1, 2, 3, 4}
	Convey("test rotate", t, func() {
		rotate(nums, k)
		So(nums, ShouldResemble, output)
	})
}
