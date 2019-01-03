package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrStr(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	output := 2
	Convey("test strStr", t, func() {
		re := strStr(haystack, needle)
		So(re, ShouldResemble, output)
	})
}
