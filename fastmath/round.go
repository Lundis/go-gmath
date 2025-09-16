package fastmath

import "math"

func Round(x float32) float32 {
	return float32(math.Round(float64(x)))
}

// RoundPos is a potentially faster Round() that only works for positive numbers.
// on AMD64 it
func RoundPos(x float32) float32 {
	const toint = float32(1 << 23)
	return (x + toint) - toint
}
