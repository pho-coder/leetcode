package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRomanToInt(t *testing.T) {
	input := "MCMXCIV"
	output := 1994
	Convey("test RomanToInt", t, func() {
		re := RomanToInt(input)
		So(re, ShouldResemble, output)
	})
}
