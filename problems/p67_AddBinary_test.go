package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddBinary(t *testing.T) {
	a := "1010"
	b := "1011"
	output := "10101"
	Convey("test addBinary", t, func() {
		re := addBinary(a, b)
		So(re, ShouldResemble, output)
	})
}
