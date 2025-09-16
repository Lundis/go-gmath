package fastmath

import "math"

type Float16 uint16

// Float32ToFloat16Fast converts float32 -> float16 bits by truncating
// mantissa (no rounding) and flushing subnormals to ±0.
// Assumes: finite input within float16 normal range (no overflow).
func Float32ToFloat16Fast(x float32) Float16 {
	u := math.Float32bits(x)
	sign := uint16(u>>16) & 0x8000
	e := (u >> 23) & 0xFF // raw exponent
	m := u & 0x7FFFFF     // mantissa (no hidden bit)

	// Normal half range: raw e in [113..142]
	if e-113 <= (142 - 113) { // branch-friendly: e >= 113 && e <= 142
		exp := uint16(e - 112)
		mant := uint16(m >> 13) // truncate
		return Float16(sign | (exp << 10) | mant)
	}

	// Everything else becomes signed zero (flush subnormals)
	return Float16(sign)
}
