package vec2

// LineIntersection calculates the intersection between lines A-B and C-D
func LineIntersection(A, B, C, D F) F {
	// https://www.gamers.org/dEngine/rsc/usenet/comp.graphics.algorithms.faq
	r_numerator := (A.Y-C.Y)*(D.X-C.X) - (A.X-C.X)*(D.Y-C.Y)
	r_denomenator := (B.X-A.X)*(D.Y-C.Y) - (B.Y-A.Y)*(D.X-C.X)
	if r_denomenator == 0.0 {
		// parallel lines
		if r_numerator == 0.0 {
			// same line
			return A
		}
		return A
	}
	r := r_numerator / r_denomenator
	return A.Plus(B.Minus(A).MulScalar(r))
}
