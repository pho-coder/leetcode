package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	input := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	output := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	Convey("test reverseString", t, func() {
		reverseString(input)
		So(input, ShouldResemble, output)
	})
}
