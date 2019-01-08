package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLengthOfLastWord(t *testing.T) {
	input1 := "Hello World"
	output := 5
	Convey("test lengthOfLastWord", t, func() {
		re := lengthOfLastWord(input1)
		So(re, ShouldResemble, output)
	})
}
