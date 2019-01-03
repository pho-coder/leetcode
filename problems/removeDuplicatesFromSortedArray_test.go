package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output := 5
	Convey("test RemoveDuplicates", t, func() {
		re := RemoveDuplicates(input)
		So(re, ShouldResemble, output)
	})
}
