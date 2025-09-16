package fastmath

import (
	"testing"
)

func BenchmarkSqrt(b *testing.B) {
	val := float32(1)
	sink := float32(0)
	for b.Loop() {
		sink += Sqrt(val)
		val += 0.1
	}
	_ = sink
}
