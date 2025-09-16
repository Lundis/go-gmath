package fastmath

import (
	"fmt"
	"math"
	"testing"
)

func TestLog2(t *testing.T) {
	for i := float32(1); i < 10000; i *= 1.5 {
		expected := float32(math.Log2(float64(i)))
		t.Run(fmt.Sprintf("Log2(%.3f) should be %.3f", i, expected), func(t *testing.T) {
			got := Log2(i)
			if math.IsNaN(float64(got)) {
				t.Errorf("Log2(%.3f) = %v, want %v", i, got, expected)
				return
			}
			diff := math.Abs(float64(got - expected))
			if diff > 0.005 {
				t.Errorf("Log2(%.3f) = %v, want %v, diff: %v", i, got, expected, diff)
			}
		})
	}
}

func BenchmarkLog2(b *testing.B) {
	val := float32(1)
	for b.Loop() {
		Log2(val)
		val += 0.1
	}
}

func BenchmarkMathLog2(b *testing.B) {
	val := float64(1)
	for b.Loop() {
		math.Log2(val)
		val += 0.1
	}
}
