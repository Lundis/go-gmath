package fastmath

import "math"

// Sqrt maps to math.Sqrt, as there is no faster sqrt than the native one
func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}
