package vec2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineIntersection(t *testing.T) {
	t.Run("perpendicular lines intersecting at origin", func(t *testing.T) {
		A := F{X: -1, Y: 0}
		B := F{X: 1, Y: 0}
		C := F{X: 0, Y: -1}
		D := F{X: 0, Y: 1}

		result := LineIntersection(A, B, C, D)
		assert.InDelta(t, 0.0, result.X, 0.001)
		assert.InDelta(t, 0.0, result.Y, 0.001)
	})

	t.Run("perpendicular lines intersecting at (5, 3)", func(t *testing.T) {
		A := F{X: 0, Y: 3}
		B := F{X: 10, Y: 3}
		C := F{X: 5, Y: 0}
		D := F{X: 5, Y: 10}

		result := LineIntersection(A, B, C, D)
		assert.InDelta(t, 5.0, result.X, 0.001)
		assert.InDelta(t, 3.0, result.Y, 0.001)
	})

	t.Run("diagonal lines intersecting", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 4, Y: 4}
		C := F{X: 0, Y: 4}
		D := F{X: 4, Y: 0}

		result := LineIntersection(A, B, C, D)
		assert.InDelta(t, 2.0, result.X, 0.001)
		assert.InDelta(t, 2.0, result.Y, 0.001)
	})

	t.Run("extended lines intersection beyond segments", func(t *testing.T) {
		// Line segments that don't intersect but their extensions do
		A := F{X: 0, Y: 0}
		B := F{X: 1, Y: 0}
		C := F{X: 2, Y: -1}
		D := F{X: 2, Y: 1}

		result := LineIntersection(A, B, C, D)
		assert.InDelta(t, 2.0, result.X, 0.001)
		assert.InDelta(t, 0.0, result.Y, 0.001)
	})

	t.Run("parallel horizontal lines", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 0}
		C := F{X: 0, Y: 5}
		D := F{X: 10, Y: 5}

		result := LineIntersection(A, B, C, D)
		// Should return A when lines are parallel
		assert.Equal(t, A.X, result.X)
		assert.Equal(t, A.Y, result.Y)
	})

	t.Run("parallel vertical lines", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 0, Y: 10}
		C := F{X: 5, Y: 0}
		D := F{X: 5, Y: 10}

		result := LineIntersection(A, B, C, D)
		// Should return A when lines are parallel
		assert.Equal(t, A.X, result.X)
		assert.Equal(t, A.Y, result.Y)
	})

	t.Run("same line (collinear)", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 10}
		C := F{X: 5, Y: 5}
		D := F{X: 15, Y: 15}

		result := LineIntersection(A, B, C, D)
		// Should return A when lines are the same
		assert.Equal(t, A.X, result.X)
		assert.Equal(t, A.Y, result.Y)
	})

	t.Run("45 degree angle intersection", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 10}
		C := F{X: 10, Y: 0}
		D := F{X: 0, Y: 10}

		result := LineIntersection(A, B, C, D)
		assert.InDelta(t, 5.0, result.X, 0.001)
		assert.InDelta(t, 5.0, result.Y, 0.001)
	})
}

func TestIntersectsLineExclusive(t *testing.T) {
	t.Run("crossing segments", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 10}
		C := F{X: 0, Y: 10}
		D := F{X: 10, Y: 0}

		assert.True(t, IntersectsLineExclusive(A, B, C, D))
		assert.True(t, IntersectsLineExclusive(C, D, A, B)) // symmetric
	})

	t.Run("perpendicular crossing segments", func(t *testing.T) {
		A := F{X: 5, Y: 0}
		B := F{X: 5, Y: 10}
		C := F{X: 0, Y: 5}
		D := F{X: 10, Y: 5}

		assert.True(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("non-intersecting parallel segments", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 0}
		C := F{X: 0, Y: 5}
		D := F{X: 10, Y: 5}

		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("non-intersecting segments (would intersect if extended)", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 1, Y: 1}
		C := F{X: 3, Y: 0}
		D := F{X: 4, Y: 1}

		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("segments sharing endpoint A-C", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 10}
		C := F{X: 0, Y: 0}
		D := F{X: 10, Y: 0}

		// Exclusive means endpoints don't count as intersection
		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("segments sharing endpoint B-D", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 5, Y: 5}
		C := F{X: 10, Y: 0}
		D := F{X: 5, Y: 5}

		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("T-junction (endpoint touches middle)", func(t *testing.T) {
		A := F{X: 0, Y: 5}
		B := F{X: 10, Y: 5}
		C := F{X: 5, Y: 0}
		D := F{X: 5, Y: 5}

		// D is on segment A-B, but it's an endpoint so should be excluded
		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("collinear non-overlapping segments", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 5, Y: 5}
		C := F{X: 10, Y: 10}
		D := F{X: 15, Y: 15}

		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("collinear overlapping segments", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 10}
		C := F{X: 5, Y: 5}
		D := F{X: 15, Y: 15}

		// Collinear overlapping segments don't "cross" in the traditional sense
		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("L-shaped segments (perpendicular but not intersecting)", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 5, Y: 0}
		C := F{X: 10, Y: 0}
		D := F{X: 10, Y: 5}

		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("crossing at very small angle", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 10, Y: 1}
		C := F{X: 0, Y: 1}
		D := F{X: 10, Y: 0}

		assert.True(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("segments at 45 degrees", func(t *testing.T) {
		A := F{X: 0, Y: 5}
		B := F{X: 10, Y: 5}
		C := F{X: 3, Y: 2}
		D := F{X: 7, Y: 8}

		assert.True(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("one segment is a point (degenerate case)", func(t *testing.T) {
		A := F{X: 5, Y: 5}
		B := F{X: 5, Y: 5}
		C := F{X: 0, Y: 0}
		D := F{X: 10, Y: 10}

		// Point on line segment
		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("both segments are points (degenerate case)", func(t *testing.T) {
		A := F{X: 5, Y: 5}
		B := F{X: 5, Y: 5}
		C := F{X: 5, Y: 5}
		D := F{X: 5, Y: 5}

		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("near miss (very close but not intersecting)", func(t *testing.T) {
		A := F{X: 0, Y: 0}
		B := F{X: 5, Y: 0}
		C := F{X: 5.001, Y: -1}
		D := F{X: 5.001, Y: 1}

		assert.False(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("crossing with negative coordinates", func(t *testing.T) {
		A := F{X: -10, Y: -10}
		B := F{X: 10, Y: 10}
		C := F{X: -10, Y: 10}
		D := F{X: 10, Y: -10}

		assert.True(t, IntersectsLineExclusive(A, B, C, D))
	})

	t.Run("very long segments crossing", func(t *testing.T) {
		A := F{X: -1000, Y: 0}
		B := F{X: 1000, Y: 0}
		C := F{X: 0, Y: -1000}
		D := F{X: 0, Y: 1000}

		assert.True(t, IntersectsLineExclusive(A, B, C, D))
	})
}

func BenchmarkLineIntersection(b *testing.B) {
	A := F{X: 0, Y: 0}
	B := F{X: 10, Y: 10}
	C := F{X: 0, Y: 10}
	D := F{X: 10, Y: 0}

	for b.Loop() {
		LineIntersection(A, B, C, D)
		A.X += 0.001
		A.Y += 0.001
		B.X += 0.001
		B.Y += 0.001
		C.Y += 0.001
		D.X += 0.001
	}
}

func BenchmarkIntersectsLineExclusive(b *testing.B) {
	A := F{X: 0, Y: 0}
	B := F{X: 10, Y: 10}
	C := F{X: 0, Y: 10}
	D := F{X: 10, Y: 0}

	for b.Loop() {
		_ = IntersectsLineExclusive(A, B, C, D)
		A.X += 0.001
		A.Y += 0.001
		B.X += 0.001
		B.Y += 0.001
		C.Y += 0.001
		D.X += 0.001
	}
}
