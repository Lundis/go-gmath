package fastmath

import (
	"math"
	"testing"
)

func BenchmarkCopySign(b *testing.B) {
	val := float32(1)
	val2 := float32(-1)
	sink := float32(0)
	for b.Loop() {
		sink += Copysign(val2, val)
		val *= -1
		val2 *= -1
	}
	_ = sink
}

func BenchmarkMathCopySign(b *testing.B) {
	val := float32(1)
	val2 := float32(-1)
	sink := float32(0)
	for b.Loop() {
		sink += float32(math.Copysign(float64(val2), float64(val)))
		val *= -1
		val2 *= -1
	}
	_ = sink
}
