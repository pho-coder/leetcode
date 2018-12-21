package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	Convey("test two sum", t, func() {
		re := TwoSum1(nums, target)
		So(re, ShouldResemble, []int{0, 1})
	})
}

func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	Convey("test two sum", t, func() {
		re := TwoSum2(nums, target)
		So(re, ShouldResemble, []int{0, 1})
	})
}

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum1(nums, target)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2(nums, target)
	}
}
