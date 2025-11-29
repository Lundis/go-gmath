package vec3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomF(t *testing.T) {
	for i := float32(0); i < 100; i++ {
		a := NewRandomF(-1, i)
		assert.True(t, -1 <= a.X && a.X <= i)
		assert.True(t, -1 <= a.Y && a.Y <= i)
		assert.True(t, -1 <= a.Z && a.Z <= i)
	}
}

func TestEquals(t *testing.T) {
	a := F{5, 3, 7}
	b := F{4, 3, 7}
	c := b

	assert.False(t, a.Equals(b))
	assert.True(t, c.Equals(b))
}

func TestString(t *testing.T) {
	a := F{3, 6, 9}
	assert.Equal(t, "(3, 6, 9)", a.String())
	c := F{3.5, 6.3, 7.44326}
	assert.Equal(t, "(3.5000, 6.3000, 7.4433)", c.String())
}

func TestIsZero(t *testing.T) {
	assert.True(t, F{0, 0, 0}.IsZero())
	assert.False(t, F{0.000001, 0.000001, 0.000001}.IsZero())
	assert.False(t, F{1, 0, 0}.IsZero())
	assert.False(t, F{0, 1, 0}.IsZero())
	assert.False(t, F{0, 0, 1}.IsZero())
}

func TestPlus(t *testing.T) {
	a := F{1, 2, 5}
	b := F{4, 3, 1}
	res := a.Plus(b)

	assert.Equal(t, float32(5), res.X)
	assert.Equal(t, float32(5), res.Y)
	assert.Equal(t, float32(6), res.Z)
}

func TestAddScalar(t *testing.T) {
	a := F{1, 2, 3}
	res := a.AddScalar(1)

	assert.Equal(t, float32(2), res.X)
	assert.Equal(t, float32(3), res.Y)
	assert.Equal(t, float32(4), res.Z)
}

func TestAddScalars(t *testing.T) {
	a := F{1, 2, 3}
	res := a.AddScalars(1, 2, 3)

	assert.Equal(t, float32(2), res.X)
	assert.Equal(t, float32(4), res.Y)
	assert.Equal(t, float32(6), res.Z)
}

func TestMinus(t *testing.T) {
	a := F{1, 2, 3}
	b := F{4, 3, 2}
	res := a.Minus(b)

	assert.Equal(t, float32(-3), res.X)
	assert.Equal(t, float32(-1), res.Y)
	assert.Equal(t, float32(1), res.Z)
}

func TestSubScalar(t *testing.T) {
	a := F{1, 2, 3}
	res := a.SubScalar(1)

	assert.Equal(t, float32(0), res.X)
	assert.Equal(t, float32(1), res.Y)
	assert.Equal(t, float32(2), res.Z)
}

func TestSubScalars(t *testing.T) {
	a := F{1, 2, 3}
	res := a.SubScalars(1, 2, 3)

	assert.Equal(t, float32(0), res.X)
	assert.Equal(t, float32(0), res.Y)
	assert.Equal(t, float32(0), res.Z)
}

func TestMul(t *testing.T) {
	a := F{1, 2, 3}
	b := F{4, 3, 7}
	res := a.Mul(b)

	assert.Equal(t, float32(4), res.X)
	assert.Equal(t, float32(6), res.Y)
	assert.Equal(t, float32(21), res.Z)
}

func TestMulScalar(t *testing.T) {
	a := F{1, 2, 3}
	const b = 4
	res := a.MulScalar(b)

	assert.Equal(t, float32(4), res.X)
	assert.Equal(t, float32(8), res.Y)
	assert.Equal(t, float32(12), res.Z)
}

func TestMulScalars(t *testing.T) {
	a := F{1, 2, 3}
	res := a.MulScalars(1, 2, 3)

	assert.Equal(t, float32(1), res.X)
	assert.Equal(t, float32(4), res.Y)
	assert.Equal(t, float32(9), res.Z)
}

func TestDivMethod(t *testing.T) {
	a := F{1, 2, 3}
	b := F{4, 3, 2}
	res := a.Div(b)

	assert.Equal(t, float32(0.25), res.X)
	assert.Equal(t, float32(0.6666666666666666), res.Y)
	assert.Equal(t, float32(1.5), res.Z)
}

func TestDivScalar(t *testing.T) {
	a := F{1, 2, 3}
	const b = 4.0
	res := a.DivScalar(b)

	assert.Equal(t, float32(0.25), res.X)
	assert.Equal(t, float32(0.5), res.Y)
	assert.Equal(t, float32(0.75), res.Z)
}

func TestDivScalars(t *testing.T) {
	a := F{1, 2, 3}
	res := a.DivScalars(1.0, 2, 3)

	assert.Equal(t, float32(1), res.X)
	assert.Equal(t, float32(1), res.Y)
	assert.Equal(t, float32(1), res.Z)
}

func TestMagnitude(t *testing.T) {
	a := F{3, 2, 1}
	res := a.Magnitude()

	assert.Equal(t, float32(3.74165738677), res)
}

func TestNormalized(t *testing.T) {
	a := F{3, 2, 1}
	b := F{0.000003, 0.000002, 0.000001}
	aa := a.Normalized().MulScalar(1000).Round()
	bb := b.Normalized().MulScalar(1000).Round()

	assert.True(t, aa.Equals(bb))
}

func TestNormalizeZero(t *testing.T) {
	a := F{0, 0, 0}
	aa := a.Normalized()

	assert.True(t, aa.Equals(a))
}

func TestAbs(t *testing.T) {
	a := F{-1, -2, -3}
	res := a.Abs()
	assert.Equal(t, float32(1), res.X)
	assert.Equal(t, float32(2), res.Y)
	assert.Equal(t, float32(3), res.Z)
}

func TestClamp(t *testing.T) {
	low := F{-5, -13, -7}
	high := F{17, 10, 13}
	a := F{-100, -100, -100}.Clamp(low, high)
	b := F{100, 100, 100}.Clamp(low, high)
	assert.True(t, low.Equals(a))
	assert.True(t, high.Equals(b))
}

func TestRound(t *testing.T) {
	a := F{-3.4, 5.5, 6.6}.Round()
	assert.Equal(t, float32(-3), a.X)
	assert.Equal(t, float32(6), a.Y)
	assert.Equal(t, float32(7), a.Z)
}

func TestCeil(t *testing.T) {
	a := F{-3.4, 5.5, 6.6}.Ceil()
	assert.Equal(t, float32(-3), a.X)
	assert.Equal(t, float32(6), a.Y)
	assert.Equal(t, float32(7), a.Z)
}

func TestFloor(t *testing.T) {
	a := F{-3.4, 5.5, 6.6}.Floor()
	assert.Equal(t, float32(-3), a.X)
	assert.Equal(t, float32(5), a.Y)
	assert.Equal(t, float32(6), a.Z)
}

func TestComponents(t *testing.T) {
	x, y, z := F{-3.4, 5.5, 6.6}.Components()
	assert.Equal(t, float32(-3.4), x)
	assert.Equal(t, float32(5.5), y)
	assert.Equal(t, float32(6.6), z)
}

func TestMinMax(t *testing.T) {
	a := F{-5, 10, 23}
	b := F{17, -13, -7}
	c := a.Min(b)
	assert.Equal(t, float32(-5), c.X)
	assert.Equal(t, float32(-13), c.Y)
	assert.Equal(t, float32(-7), c.Z)
	c = a.Max(b)
	assert.Equal(t, float32(17), c.X)
	assert.Equal(t, float32(10), c.Y)
	assert.Equal(t, float32(23), c.Z)
}

func TestIsBetweenInclusive(t *testing.T) {
	a := F{1, 2, 7}
	b := F{3, 6, 10}
	c := F{5, 10, 13}

	assert.Equal(t, false, a.IsBetween(b, c))
	assert.Equal(t, true, b.IsBetween(a, c))
	assert.Equal(t, false, c.IsBetween(a, b))
}

func TestDot(t *testing.T) {
	a := F{1, 2, -6}
	b := F{4, 3, 7}
	res := a.Dot(b)

	assert.Equal(t, res, float32(4+6-42))
}

func TestDistanceToLine(t *testing.T) {
	/*
				 |           p2
		         |
				 b
		p3       |
				 a     p

	*/
	a := F{0, 0, 0}
	b := F{0, 1, 0}

	p := F{1, 0, 0}
	assert.Equal(t, float32(1), p.DistanceToLine(a, b))

	p2 := F{2, 2, 0}
	assert.Equal(t, float32(2), p2.DistanceToLine(a, b))

	p3 := F{-0.5, 0.5, 0}
	assert.Equal(t, float32(0.5), p3.DistanceToLine(a, b))
}
