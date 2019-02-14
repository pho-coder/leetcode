package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMoveZeroes(t *testing.T) {
	input := []int{0, 1, 0, 3, 0, 12}
	output := []int{1, 3, 12, 0, 0, 0}
	Convey("test moveZeroes", t, func() {
		moveZeroes(input)
		So(input, ShouldResemble, output)
	})
}
