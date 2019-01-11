package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClimbStairs(t *testing.T) {
	input := 6
	output := 12
	Convey("test climbStairs", t, func() {
		re := climbStairs(input)
		So(re, ShouldResemble, output)
	})
}
