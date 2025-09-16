package fastmath

import "math"

func Mod(a, b float32) float32 {
	div := float32(int32(a / b))
	result := a - div*b
	if result < 0 && a > 0 && b > 0 { // handle rounding errors for large values
		result += b
	}
	return result
}

// ModAbs is like Mod, but handles cases like (-0.5 % 2) == 1.5
func ModAbs(a, b float32) float32 {
	div := float32(int32(a / b))
	if a < 0 {
		// instead of being in the negative range (-b, 0),
		// this moves us to the expected (0, b)
		div -= 1
	}
	result := a - div*b
	if result < 0 { // handle rounding errors for large values
		result += b
	}
	return result
}

// Modf splits x into integer and fractional parts (float32).
// Assumes x is finite (no ±Inf/NaN).
// - intPart is x truncated toward zero.
// - fracPart = x - intPart, with the same sign as x (incl. signed zero).
func Modf(x float32) (intPart, fracPart float32) {
	bits := math.Float32bits(x)
	exp := int((bits >> 23) & 0xFF) // biased exponent (bias=127)

	// If |x| >= 2^23: no fractional bits are representable → frac is signed zero.
	if exp >= 150 { // 127 + 23
		intPart = x
		fracPart = math.Float32frombits(bits & 0x80000000) // ±0 with x's sign
		return
	}

	// If |x| < 1: integer part is signed zero, fractional is x.
	if exp < 127 {
		intPart = math.Float32frombits(bits & 0x80000000) // ±0
		fracPart = x
		return
	}

	// 1 <= |x| < 2^23: clear the fractional mantissa bits.
	shift := uint32(150 - exp)              // number of fractional mantissa bits
	mask := (uint32(1) << shift) - 1        // low 'shift' bits
	intBits := bits &^ mask                 // zero out fractional bits
	intPart = math.Float32frombits(intBits) // truncated toward zero
	fracPart = x - intPart                  // same sign as x
	return
}
