package vec2

// LineIntersection calculates the (extended) intersection between lines A-B and C-D
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

// IntersectsLineExclusive returns true if line segments A-B and C-D intersect (excluding endpoints)
func IntersectsLineExclusive(A, B, C, D F) bool {
	// Calculate cross products to determine orientation
	// If A-B and C-D intersect, then C and D must be on opposite sides of A-B,
	// and A and B must be on opposite sides of C-D

	// Cross product of (B-A) × (C-A)
	cross1 := (B.X-A.X)*(C.Y-A.Y) - (B.Y-A.Y)*(C.X-A.X)
	// Cross product of (B-A) × (D-A)
	cross2 := (B.X-A.X)*(D.Y-A.Y) - (B.Y-A.Y)*(D.X-A.X)

	if cross1*cross2 >= 0 {
		return false
	}

	// Cross product of (D-C) × (A-C)
	cross3 := (D.X-C.X)*(A.Y-C.Y) - (D.Y-C.Y)*(A.X-C.X)
	// Cross product of (D-C) × (B-C)
	cross4 := (D.X-C.X)*(B.Y-C.Y) - (D.Y-C.Y)*(B.X-C.X)

	return cross3*cross4 < 0
}
