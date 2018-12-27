package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPalindrome(t *testing.T) {
	input := 121
	output := true
	Convey("test IsPalindrome", t, func() {
		re := IsPalindrome(input)
		So(re, ShouldResemble, output)
	})
}
