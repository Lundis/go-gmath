package fastmath

import "math"

func Equald(a, b, precision float64) bool {
	return math.Abs(a-b) <= precision
}

func Equalf(a, b, precision float32) bool {
	return math.Abs(float64(a-b)) <= float64(precision)
}
