package fastmath

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	for x := float32(-2); x < 2; x += 0.01 {
		actual := Round(x)
		expected := math.Round(float64(x))
		diff := float64(actual) - expected
		if math.Abs(diff) > 0.004 {
			t.Errorf("Round(%f) = %f instead of %f. diff %f", x, actual, expected, diff)
		}
	}
}

func TestRoundPos(t *testing.T) {
	for x := float32(0); x < 2; x += 0.01 {
		actual := RoundPos(x)
		expected := math.Round(float64(x))
		diff := float64(actual) - expected
		if math.Abs(diff) > 0.004 {
			t.Errorf("Round(%f) = %f instead of %f. diff %f", x, actual, expected, diff)
		}
	}
}

func BenchmarkRound(b *testing.B) {
	val := float32(1)
	sink := float32(0)
	for b.Loop() {
		sink += Round(val)
		val += 0.1
	}
	_ = sink
}

func BenchmarkRoundPos(b *testing.B) {
	val := float32(1)
	sink := float32(0)
	for b.Loop() {
		sink += RoundPos(val)
		val += 0.1
	}
	_ = sink
}

func BenchmarkMathRound(b *testing.B) {
	val := float32(1)
	sink := float32(0)
	for b.Loop() {
		sink += float32(math.Round(float64(val)))
		val += 0.1
	}
	_ = sink
}
