package fastmath

import (
	"math"
	"testing"
)

func BenchmarkMathAbs(b *testing.B) {
	val := float32(1)
	sink := float32(0)
	for b.Loop() {
		sink += float32(math.Abs(float64(val)))
		val *= -1
	}
	_ = sink
}
