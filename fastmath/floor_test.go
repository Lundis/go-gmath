package fastmath

import (
	"math"
	"testing"
)

func BenchmarkMathFloor(b *testing.B) {
	val := float32(1)
	sink := float32(0)
	for b.Loop() {
		sink += float32(math.Floor(float64(val)))
		val += 0.01
	}
	_ = sink
}
