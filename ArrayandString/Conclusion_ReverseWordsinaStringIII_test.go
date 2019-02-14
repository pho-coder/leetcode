package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseWords(t *testing.T) {
	input := "Let's take LeetCode contest"
	output := "s'teL ekat edoCteeL tsetnoc"
	Convey("test reverseWords", t, func() {
		re := reverseWords(input)
		So(re, ShouldResemble, output)
	})
}
