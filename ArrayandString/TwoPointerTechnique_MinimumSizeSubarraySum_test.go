package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinSubArrayLen(t *testing.T) {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	output := 2
	Convey("test minSubArrayLen", t, func() {
		re := minSubArrayLen(s, nums)
		So(re, ShouldResemble, output)
	})
}
