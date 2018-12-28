package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestCommonPrefix(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	output := "fl"
	Convey("test LongestCommonPrefix", t, func() {
		re := LongestCommonPrefix(input)
		So(re, ShouldResemble, output)
	})
}
