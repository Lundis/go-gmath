package fastmath

import (
	"fmt"
	"math"
	"testing"
)

func TestMod(t *testing.T) {
	tests := []struct {
		a, b float32
		want float32
	}{
		{0, 2, 0},
		{1, 2, 1},
		{2, 2, 0},
		{2.5, 2, 0.5},
		{0.5, 2, 0.5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %% %v should be %v", tt.a, tt.b, tt.want), func(t *testing.T) {
			if got := Mod(tt.a, tt.b); got != tt.want {
				t.Errorf("Mod(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestModAbs(t *testing.T) {
	tests := []struct {
		a, b float32
		want float32
	}{
		{0, 2, 0},
		{1, 2, 1},
		{2, 2, 0},
		{0.5, 2, 0.5},
		{-0.5, 2, 1.5},
		{-3, 2, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %% %v should be %v", tt.a, tt.b, tt.want), func(t *testing.T) {
			if got := ModAbs(tt.a, tt.b); got != tt.want {
				t.Errorf("Mod(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestModf(t *testing.T) {
	for i := float32(0); i < 100; i += 0.01 {
		actualInt, actualFrac := Modf(i)
		expectedInt, expectedFrac := math.Modf(float64(i))
		if actualInt != float32(expectedInt) {
			t.Errorf("Modf(%f) int %f, want %f", i, actualInt, expectedInt)
		}
		diff := float64(actualFrac) - expectedFrac
		if math.Abs(diff) > 0.0001 {
			t.Errorf("Modf(%f) frac %f, want %f diff %f",
				i, actualFrac, float32(expectedFrac), diff)
		}
	}
	// this specific value triggers rounding errors in the floor logic
	Cos(412032.44)
}

func BenchmarkMod(b *testing.B) {
	val := float32(1)
	for b.Loop() {
		Mod(val, math.Pi)
		val += 0.1
	}
}

func BenchmarkModAbs(b *testing.B) {
	val := float32(1)
	for b.Loop() {
		ModAbs(val, math.Pi)
		val += 0.1
	}
}

func BenchmarkMathMod(b *testing.B) {
	val := float64(1)
	for b.Loop() {
		math.Mod(val, math.Pi)
		val += 0.1
	}
}

func BenchmarkModf(b *testing.B) {
	val := float32(1)
	for b.Loop() {
		Modf(val)
		val += 0.1
	}
}

func BenchmarkMathModf(b *testing.B) {
	val := float64(1)
	for b.Loop() {
		math.Modf(val)
		val += 0.1
	}
}
