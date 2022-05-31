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
