package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArrayPairSum(t *testing.T) {
	input := []int{1, 4, 3, 2}
	output := 4
	Convey("test arrayPairSum", t, func() {
		re := arrayPairSum(input)
		So(re, ShouldResemble, output)
	})
}
