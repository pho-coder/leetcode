package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMySqrt(t *testing.T) {
	input := 24
	output := 4
	Convey("test mySqrt", t, func() {
		re := mySqrt(input)
		So(re, ShouldResemble, output)
	})
}
