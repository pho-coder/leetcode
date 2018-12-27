package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverse(t *testing.T) {
	input := 120
	output := 21
	Convey("test Reverse", t, func() {
		re := Reverse(input)
		So(re, ShouldResemble, output)
	})
}
