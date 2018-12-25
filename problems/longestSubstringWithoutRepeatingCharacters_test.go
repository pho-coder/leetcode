package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLengthOfLongestSubstring1(t *testing.T) {
	aStr := "abba"
	Convey("test LengthOfLongestSubstring", t, func() {
		So(LengthOfLongestSubstring1(aStr), ShouldEqual, 2)
	})
}
