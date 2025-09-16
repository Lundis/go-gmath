package fastmath

import "math"

func Log2(val float32) float32 {
	// Precondition: val > 0. For val <= 0, use math.Log2 for correct IEEE behavior.
	x := math.Float32bits(val)

	// Exponent term (as a number, not a bit pattern).
	e := (x >> 23) & 0xFF
	log2 := float32(e) - 128.0 // matches the polynomial's ~+1 bias

	// Build mantissa in [1, 2)
	x = (x & ^uint32(0xFF<<23)) | (127 << 23)
	u := math.Float32frombits(x)

	// Polynomial approximation for log2(u) + 1 (so +1 is compensated by -128 above).
	log2 += ((-0.34484843)*u+2.02466578)*u - 0.67487759
	return log2
}
