package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRow(t *testing.T) {
	input := 3
	output := []int{1, 3, 3, 1}
	Convey("test getRow", t, func() {
		re := getRow(input)
		So(re, ShouldResemble, output)
	})
}
