package vec2

import "github.com/Lundis/go-gmath/fastmath"

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
	return A.Add(B.Sub(A).MulScalar(r))
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

// IntersectsInfiniteLineCircle returns true if an infinite line segment A-B intersects circle with given center and radius
// Based on https://mathworld.wolfram.com/Circle-LineIntersection.html
func IntersectsInfiniteLineCircle(A, B, center F, radius float32) bool {
	// Translate line segment to origin
	Ax := A.X - center.X
	Ay := A.Y - center.Y
	Bx := B.X - center.X
	By := B.Y - center.Y

	// Line segment vector
	dx := Bx - Ax
	dy := By - Ay

	dr := fastmath.Sqrt(dx*dx + dy*dy)
	D := Ax*By - Bx*Ay

	discriminant := radius*radius*dr*dr - D*D
	return discriminant >= 0
}

func IntersectsLineCircleInclusive(A, B, center F, radius float32) bool {
	// Translate line segment to origin (relative to circle center)
	Ax := A.X - center.X
	Ay := A.Y - center.Y
	Bx := B.X - center.X
	By := B.Y - center.Y

	radiusSq := radius * radius

	// Check if either endpoint is inside or on the circle
	distASq := Ax*Ax + Ay*Ay
	distBSq := Bx*Bx + By*By
	if distASq <= radiusSq || distBSq <= radiusSq {
		return true
	}

	// Line segment vector
	dx := Bx - Ax
	dy := By - Ay

	// Find closest point on line segment to circle center (at origin)
	// Parameter t represents position along line segment (0 = A, 1 = B)
	t := -(Ax*dx + Ay*dy) / (dx*dx + dy*dy)

	// closest point is outside segment
	if t < 0 || t > 1 {
		return false
	}

	// Calculate closest point on segment
	closestX := Ax + t*dx
	closestY := Ay + t*dy

	// Check if closest point is within or on the circle
	distSq := closestX*closestX + closestY*closestY
	return distSq <= radiusSq
}

// ClosestPointOnLineSegmentF returns the point on the line a-b closest to target.
// t is the ratio between a and b of that point
func ClosestPointOnLineSegmentF(a, b, target F) (closest F, t float32) {
	ab := b.Sub(a)
	at := target.Sub(a)
	t = at.Dot(ab) / ab.Dot(ab)
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}
	closest = a.Add(ab.MulScalar(t))
	return
}

// ClosestPointOnLineSegmentD returns the point on the line a-b closest to target.
// t is the ratio between a and b of that point
func ClosestPointOnLineSegmentD(a, b, target D) (closest D, t float64) {
	ab := b.Sub(a)
	at := target.Sub(a)
	t = at.Dot(ab) / ab.Dot(ab)
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}
	closest = a.Add(ab.MulScalar(t))
	return
}

// ClosestPointOnLineF returns the point on the infinite line a-b closest to target
func ClosestPointOnLineF(a, b, target F) (closest F) {
	ab := b.Sub(a)
	at := target.Sub(a)
	t := at.Dot(ab) / ab.Dot(ab)
	closest = a.Add(ab.MulScalar(t))
	return
}

// ClosestPointOnLineD returns the point on the infinite line a-b closest to target
func ClosestPointOnLineD(a, b, target D) (closest D) {
	ab := b.Sub(a)
	at := target.Sub(a)
	t := at.Dot(ab) / ab.Dot(ab)
	closest = a.Add(ab.MulScalar(t))
	return
}
