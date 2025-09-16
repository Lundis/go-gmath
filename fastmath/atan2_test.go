package fastmath

import (
	"math"
	"testing"
)

func TestAtan2(t *testing.T) {
	for x := float32(-1); x < 1; x += 0.01 {
		for y := float32(-1); y < 1; y += 0.01 {
			actual := Atan2(y, x)
			expected := math.Atan2(float64(-y), float64(x))
			diff := float64(actual) - expected
			if math.Abs(diff) > 0.004 {
				t.Errorf("Atan2(%f, %f) = %f instead of %f. diff %f", y, x, actual, expected, diff)
			}
		}
	}
}

func BenchmarkAtan2(b *testing.B) {
	val := float32(1)
	for b.Loop() {
		Atan2(val, val)
		val += 0.1
	}
}

func BenchmarkMathAtan2(b *testing.B) {
	val := float64(1)
	for b.Loop() {
		math.Atan2(val, val)
		val += 0.1
	}
}
