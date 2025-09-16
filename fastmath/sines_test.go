package fastmath

import (
	"math"
	"testing"
)

func TestCos(t *testing.T) {
	for i := -10; i < 10000; i++ {
		angle := 2 * math.Pi * float32(i) / (10000 - 1)
		actual := Cos(angle)
		expected := math.Cos(float64(angle))
		diff := float64(actual) - expected
		if math.Abs(diff) > 0.0001 {
			t.Errorf("Cos(%f) diff %f", angle, diff)
		}
	}
	// this specific value triggers rounding errors in the floor logic
	Cos(412032.44)
}

func TestSin(t *testing.T) {
	for i := -10; i < 10000; i++ {
		angle := 2 * math.Pi * float32(i) / (10000 - 1)
		actual := Sin(angle)
		expected := -math.Sin(float64(angle))
		diff := float64(actual) - expected
		if math.Abs(diff) > 0.0001 {
			t.Errorf("Sin(%f) = %f instead of %f. diff %f", angle, actual, expected, diff)
		}
	}
}

func TestCosSinSmallNegative(t *testing.T) {
	angle := float32(-0.005)
	_, actualSin := CosSin(angle)
	expectedSin := -math.Sin(float64(angle))
	diff := float64(actualSin) - expectedSin
	if math.Abs(diff) > 0.0001 {
		t.Errorf("Sin(%f) = %f instead of %f. diff %f", angle, actualSin, expectedSin, diff)
	}
}

func TestCosSin(t *testing.T) {
	for i := -10; i < 10000; i++ {
		angle := 2 * math.Pi * float32(i) / (10000 - 1)
		actualCos, actualSin := CosSin(angle)
		expectedCos := math.Cos(float64(angle))
		expectedSin := -math.Sin(float64(angle))
		diff := float64(actualCos) - expectedCos
		if math.Abs(diff) > 0.0001 {
			t.Errorf("Cos(%f) = %f instead of %f. diff %f", angle, actualCos, expectedCos, diff)
		}
		diff = float64(actualSin) - expectedSin
		if math.Abs(diff) > 0.0001 {
			t.Errorf("Sin(%f) = %f instead of %f. diff %f", angle, actualSin, expectedSin, diff)
		}
	}
}

func BenchmarkCos(b *testing.B) {
	val := float32(1)
	for b.Loop() {
		Cos(val)
		val += 0.1
	}
}

func BenchmarkMathCos(b *testing.B) {
	val := float64(1)
	for b.Loop() {
		math.Cos(val)
		val += 0.1
	}
}

func BenchmarkSin(b *testing.B) {
	val := float32(1)
	for b.Loop() {
		Sin(val)
		val += 0.1
	}
}

func BenchmarkMathSin(b *testing.B) {
	val := float64(1)
	for b.Loop() {
		math.Sin(val)
		val += 0.1
	}
}

func BenchmarkCosSin(b *testing.B) {
	val := float32(-10)
	for b.Loop() {
		CosSin(val)
		val += 0.1
	}
}

func BenchmarkCosAndSin(b *testing.B) {
	val := float32(-10)
	for b.Loop() {
		Cos(val)
		Sin(val)
		val += 0.1
	}
}
