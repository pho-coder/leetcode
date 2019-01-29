package ArrayandString

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerate(t *testing.T) {
	input := 5
	output := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	Convey("test generate", t, func() {
		re := generate(input)
		So(re, ShouldResemble, output)
	})
}
