package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	output := []int{1, 2}
	Convey("test twoSum", t, func() {
		re := twoSum(numbers, target)
		So(re, ShouldResemble, output)
	})
}
