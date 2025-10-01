package fastmath

import "math"

// Copysign returns a value with the value of f and the sign of sign.
// According to benchmarks, it's faster to use math.CopySign on WASM,
// so this is here purely for utility for non-critical code.
func Copysign(f, sign float32) float32 {
	const signBit = 1 << 31
	return math.Float32frombits(math.Float32bits(f)&^signBit | math.Float32bits(sign)&signBit)
}
