package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	input := []int{1, 1, 0, 1, 1, 1}
	output := 3
	Convey("test findMaxConsecutiveOnes", t, func() {
		re := findMaxConsecutiveOnes(input)
		So(re, ShouldResemble, output)
	})
}
