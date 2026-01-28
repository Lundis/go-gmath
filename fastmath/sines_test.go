package fastmath

import (
	"math"
	"testing"
)

const cosSinPrecisionRequired = 0.0001
const fastCosSinPrecisionRequired = 0.0004

func TestCos(t *testing.T) {
	for i := -10; i < 10000; i++ {
		angle := 2 * math.Pi * float32(i) / (10000 - 1)
		actual := Cos(angle)
		expected := math.Cos(float64(angle))
		diff := float64(actual) - expected
		if math.Abs(diff) > cosSinPrecisionRequired {
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
		if math.Abs(diff) > cosSinPrecisionRequired {
			t.Errorf("Sin(%f) = %f instead of %f. diff %f", angle, actual, expected, diff)
		}
	}
}

func TestCosSinSmallNegative(t *testing.T) {
	angle := float32(-0.005)
	_, actualSin := CosSin(angle)
	expectedSin := -math.Sin(float64(angle))
	diff := float64(actualSin) - expectedSin
	if math.Abs(diff) > cosSinPrecisionRequired {
		t.Errorf("Sin(%f) = %f instead of %f. diff %f", angle, actualSin, expectedSin, diff)
	}
}

func TestCosSin(t *testing.T) {
	explicitTestValues := []struct{ angle, expectedCos, expectedSin float32 }{
		{0, 1, 0},
		{math.Pi / 2, 0, -1},
		{-math.Pi / 2, 0, 1},
		{math.Pi, -1, 0},
		{-math.Pi, -1, 0},
	}
	for _, test := range explicitTestValues {
		actualCos, actualSin := CosSin(test.angle)
		diff := float64(actualCos) - float64(test.expectedCos)
		if math.Abs(diff) > cosSinPrecisionRequired {
			t.Errorf("Cos(%f) = %f instead of %f. diff %f", test.angle, actualCos, test.expectedCos, diff)
		}
		diff = float64(actualSin) - float64(test.expectedSin)
		if math.Abs(diff) > cosSinPrecisionRequired {
			t.Errorf("Sin(%f) = %f instead of %f. diff %f", test.angle, actualSin, test.expectedSin, diff)
		}
	}
	for i := -10; i < 10000; i++ {
		angle := 2 * math.Pi * float32(i) / (10000 - 1)
		actualCos, actualSin := CosSin(angle)
		expectedCos := math.Cos(float64(angle))
		expectedSin := -math.Sin(float64(angle))
		diff := float64(actualCos) - expectedCos
		if math.Abs(diff) > cosSinPrecisionRequired {
			t.Errorf("Cos(%f) = %f instead of %f. diff %f", angle, actualCos, expectedCos, diff)
		}
		diff = float64(actualSin) - expectedSin
		if math.Abs(diff) > cosSinPrecisionRequired {
			t.Errorf("Sin(%f) = %f instead of %f. diff %f", angle, actualSin, expectedSin, diff)
		}
	}
}

func TestCosSinFast(t *testing.T) {
	explicitTestValues := []struct{ angle, expectedCos, expectedSin float32 }{
		{0, 1, 0},
		{math.Pi / 2, 0, -1},
		{-math.Pi / 2, 0, 1},
		{math.Pi, -1, 0},
		{-math.Pi, -1, 0},
	}
	for _, test := range explicitTestValues {
		actualCos, actualSin := CosSinFast(test.angle)
		diff := float64(actualCos) - float64(test.expectedCos)
		if math.Abs(diff) > fastCosSinPrecisionRequired {
			t.Errorf("Cos(%f) = %f instead of %f. diff %f", test.angle, actualCos, test.expectedCos, diff)
		}
		diff = float64(actualSin) - float64(test.expectedSin)
		if math.Abs(diff) > fastCosSinPrecisionRequired {
			t.Errorf("Sin(%f) = %f instead of %f. diff %f", test.angle, actualSin, test.expectedSin, diff)
		}
	}
	for i := -10; i < 10000; i++ {
		angle := 2 * math.Pi * float32(i) / (10000 - 1)
		actualCos, actualSin := CosSinFast(angle)
		expectedCos := math.Cos(float64(angle))
		expectedSin := -math.Sin(float64(angle))
		diff := float64(actualCos) - expectedCos
		if math.Abs(diff) > fastCosSinPrecisionRequired {
			t.Errorf("Cos(%f) = %f instead of %f. diff %f", angle, actualCos, expectedCos, diff)
		}
		diff = float64(actualSin) - expectedSin
		if math.Abs(diff) > fastCosSinPrecisionRequired {
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
	sink := float32(0)
	for b.Loop() {
		cos, sin := CosSin(val)
		sink += cos
		sink += sin
		val += 0.1
	}
	_ = sink
}

func BenchmarkMathSinCos(b *testing.B) {
	val := float32(-10)
	sink := float32(0)
	for b.Loop() {
		sin, cos := math.Sincos(float64(val))
		sink += float32(cos)
		sink += float32(sin)

		val += 0.1
	}
	_ = sink
}

func BenchmarkCosSinFast(b *testing.B) {
	val := float32(-10)
	sink := float32(0)
	for b.Loop() {
		cos, sin := CosSinFast(val)
		sink += cos
		sink += sin
		val += 0.1
	}
	_ = sink
}
