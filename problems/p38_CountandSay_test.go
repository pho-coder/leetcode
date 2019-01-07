package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountAndSay(t *testing.T) {
	input1 := 6
	output := "312211"
	Convey("test countAndSay", t, func() {
		re := countAndSay1(input1)
		So(re, ShouldResemble, output)
	})
}
