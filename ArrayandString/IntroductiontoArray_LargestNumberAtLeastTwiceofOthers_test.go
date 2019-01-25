package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDominantIndex(t *testing.T) {
	input := []int{3, 6, 1, 0}
	output := 1
	Convey("test dominantIndex", t, func() {
		re := dominantIndex(input)
		So(re, ShouldResemble, output)
	})
}
