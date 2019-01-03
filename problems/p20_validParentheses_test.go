package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsValid(t *testing.T) {
	input := "{[]}"
	output := true
	Convey("test IsValid", t, func() {
		re := IsValid(input)
		So(re, ShouldResemble, output)
	})
}
