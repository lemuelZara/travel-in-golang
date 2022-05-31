package benchmark

import (
	"testing"

	"github.com/lemuelZara/travel-in-golang/basics"
)

func Benchmark_FindEvenNumbersByRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basics.FindEvenNumbersByRange(10, 20)
	}
}

func Benchmark_BubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basics.BubbleSort([]int{10, 9, 8, 7, 6})
	}
}
