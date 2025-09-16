package vec2

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewPolarF(t *testing.T) {
	right := NewPolarF(0, 1000).Round()

	assert.Equal(t, float32(1000), right.X)
	assert.Equal(t, float32(0), right.Y)

	up := NewPolarF(0.5*math.Pi, 1000).Round()
	assert.Equal(t, float32(0), up.X)
	assert.Equal(t, float32(-1000), up.Y)

	left := NewPolarF(1*math.Pi, 1000).Round()
	assert.Equal(t, float32(-1000), left.X)
	assert.Equal(t, float32(0), left.Y)

	down := NewPolarF(1.5*math.Pi, 1000).Round()
	assert.Equal(t, float32(0), down.X)
	assert.Equal(t, float32(1000), down.Y)
}

func TestNewRandomF(t *testing.T) {
	for i := float32(0); i < 100; i++ {
		a := NewRandomF(-1, i)
		assert.True(t, -1 <= a.X && a.X <= i)
		assert.True(t, -1 <= a.Y && a.Y <= i)
	}
}

func TestEquals(t *testing.T) {
	a := F{5, 3}
	b := F{4, 3}

	assert.False(t, a.Equals(b))
}

func TestString(t *testing.T) {
	a := F{3, 6}
	assert.Equal(t, "(3, 6)", a.String())
	c := F{3.5, 6.3}
	assert.Equal(t, "(3.5000, 6.3000)", c.String())
}

func TestIsZero(t *testing.T) {
	assert.Equal(t, F{0, 0}.IsZero(), true)
	assert.Equal(t, F{0.000001, 0.000001}.IsZero(), false)
	assert.Equal(t, F{1, 0}.IsZero(), false)
	assert.Equal(t, F{0, 1}.IsZero(), false)
}

func TestPlus(t *testing.T) {
	a := F{1, 2}
	b := F{4, 3}
	res := a.Plus(b)

	assert.Equal(t, float32(5), res.X)
	assert.Equal(t, float32(5), res.Y)
}

func TestAddScalar(t *testing.T) {
	a := F{1, 2}
	res := a.AddScalar(1)

	assert.Equal(t, float32(2), res.X)
	assert.Equal(t, float32(3), res.Y)
}

func TestAddScalars(t *testing.T) {
	a := F{1, 2}
	res := a.AddScalars(1, 2)

	assert.Equal(t, float32(2), res.X)
	assert.Equal(t, float32(4), res.Y)
}

func TestMinus(t *testing.T) {
	a := F{1, 2}
	b := F{4, 3}
	res := a.Minus(b)

	assert.Equal(t, float32(-3), res.X)
	assert.Equal(t, float32(-1), res.Y)
}

func TestSubScalar(t *testing.T) {
	a := F{1, 2}
	res := a.SubScalar(1)

	assert.Equal(t, float32(0), res.X)
	assert.Equal(t, float32(1), res.Y)
}

func TestSubScalars(t *testing.T) {
	a := F{1, 2}
	res := a.SubScalars(1, 2)

	assert.Equal(t, float32(0), res.X)
	assert.Equal(t, float32(0), res.Y)
}

func TestMulMethod(t *testing.T) {
	a := F{1, 2}
	b := F{4, 3}
	res := a.Mul(b)

	assert.Equal(t, float32(4), res.X)
	assert.Equal(t, float32(6), res.Y)
}

func TestMulScalar(t *testing.T) {
	a := F{1, 2}
	const b = 4.0
	res := a.MulScalar(b)

	assert.Equal(t, float32(4), res.X)
	assert.Equal(t, float32(8), res.Y)
}

func TestMulScalars(t *testing.T) {
	a := F{1, 2}
	res := a.MulScalars(1.0, 2)

	assert.Equal(t, float32(1), res.X)
	assert.Equal(t, float32(4), res.Y)
}

func TestDivMethod(t *testing.T) {
	a := F{1, 2}
	b := F{4, 3}
	res := a.Div(b)

	assert.Equal(t, float32(0.25), res.X)
	assert.Equal(t, float32(0.6666666666666666), res.Y)
}

func TestDivScalar(t *testing.T) {
	a := F{1, 2}
	const b = 4.0
	res := a.DivScalar(b)

	assert.Equal(t, float32(0.25), res.X)
	assert.Equal(t, float32(0.5), res.Y)
}

func TestDivScalars(t *testing.T) {
	a := F{1, 2}
	res := a.DivScalars(1.0, 2)

	assert.Equal(t, float32(1), res.X)
	assert.Equal(t, float32(1), res.Y)
}

func TestMagnitude(t *testing.T) {
	a := F{3, 2}
	res := a.Magnitude()

	assert.Equal(t, float32(3.605551275463989), res)
}

func TestNormalized(t *testing.T) {
	a := F{3, 2}
	b := F{0.000003, 0.000002}
	aa := a.Normalized()
	bb := b.Normalized()

	assert.True(t, aa.Equals(bb))
}

func TestNormalizeZero(t *testing.T) {
	a := F{0, 0}
	aa := a.Normalized()

	assert.True(t, aa.Equals(a))
}

func TestAngle(t *testing.T) {

	a := NewPolarF(1, 17)
	res := a.Angle()
	assert.LessOrEqual(t, math.Abs(1-float64(res)), 0.01)

	a = NewPolarF(1.84, 17)
	res = a.Angle()
	assert.LessOrEqual(t, math.Abs(1.84-float64(res)), 0.01)
}

func TestAbs(t *testing.T) {
	a := F{-1, -2}
	res := a.Abs()
	assert.Equal(t, float32(1), res.X)
	assert.Equal(t, float32(2), res.Y)
}

func TestClamp(t *testing.T) {
	low := F{-5, -13}
	high := F{17, 10}
	a := F{-100, -100}.Clamp(low, high)
	b := F{100, 100}.Clamp(low, high)
	assert.Equal(t, low.X, a.X)
	assert.Equal(t, low.Y, a.Y)
	assert.Equal(t, high.X, b.X)
	assert.Equal(t, high.Y, b.Y)
}

func TestRound(t *testing.T) {
	a := F{-3.4, 5.5}.Round()
	assert.Equal(t, float32(-3), a.X)
	assert.Equal(t, float32(6), a.Y)
}

func TestCeil(t *testing.T) {
	a := F{-3.4, 5.5}.Ceil()
	assert.Equal(t, float32(-3), a.X)
	assert.Equal(t, float32(6), a.Y)
}

func TestFloor(t *testing.T) {
	a := F{-3.4, 5.5}.Floor()
	assert.Equal(t, float32(-3), a.X)
	assert.Equal(t, float32(5), a.Y)
}

func TestSwap(t *testing.T) {
	a := F{-3.4, 5.5}.Swap()
	assert.Equal(t, float32(-3.4), a.Y)
	assert.Equal(t, float32(5.5), a.X)
}

func TestPerpendicular(t *testing.T) {
	a := F{-3.4, 5.5}.Perpendicular()
	assert.Equal(t, float32(3.4), a.Y)
	assert.Equal(t, float32(5.5), a.X)
}

func TestWithX(t *testing.T) {
	a := F{-3.4, 5.5}.WithX(478.73)
	assert.Equal(t, float32(478.73), a.X)
	assert.Equal(t, float32(5.5), a.Y)
}

func TestWithY(t *testing.T) {
	a := F{-3.4, 5.5}.WithY(478.73)
	assert.Equal(t, float32(-3.4), a.X)
	assert.Equal(t, float32(478.73), a.Y)
}

func TestComponents(t *testing.T) {
	x, y := F{-3.4, 5.5}.Components()
	assert.Equal(t, float32(-3.4), x)
	assert.Equal(t, float32(5.5), y)
}

func TestMin(t *testing.T) {
	a := F{-5, 10}
	b := F{17, -13}
	c := a.Min(b)
	assert.Equal(t, float32(-5), c.X)
	assert.Equal(t, float32(-13), c.Y)
}

func TestMax(t *testing.T) {
	a := F{-5, 10}
	b := F{17, -13}
	c := a.Max(b)
	assert.Equal(t, float32(17), c.X)
	assert.Equal(t, float32(10), c.Y)
}

func TestIsBetweenInclusive(t *testing.T) {
	a := F{1, 2}
	b := F{3, 6}
	c := F{5, 10}

	assert.Equal(t, false, a.IsBetween(b, c))
	assert.Equal(t, true, b.IsBetween(a, c))
	assert.Equal(t, false, c.IsBetween(a, b))
}

func TestDot(t *testing.T) {
	a := F{1, 2}
	b := F{4, 3}
	res := a.Dot(b)

	assert.Equal(t, res, float32(10))
}

func TestReflect(t *testing.T) {
	a := F{2, 1}
	b := F{6, 6}
	res := a.Reflect(b)

	assert.Equal(t, res.X, float32(-66))
	assert.Equal(t, res.Y, float32(-30))
}

func TestDistanceToLine(t *testing.T) {
	/*
				 |           p2
		         |
				 b
		p3       |
				 a     p

	*/
	a := F{0, 0}
	b := F{0, 1}

	p := F{1, 0}
	assert.Equal(t, float32(1), p.DistanceToLine(a, b))

	p2 := F{2, 2}
	assert.Equal(t, float32(2), p2.DistanceToLine(a, b))

	p3 := F{-0.5, 0.5}
	assert.Equal(t, float32(0.5), p3.DistanceToLine(a, b))
}

func TestSideOfLine(t *testing.T) {
	/*
				 |           p2
		         |
				 b
		p3       |
				 a     p
	*/
	a := F{0, 0}
	b := F{0, 1}

	p := F{1, 0}
	assert.Less(t, p.SideOfLine(a, b), float32(0))

	p2 := F{2, 2}
	assert.Less(t, p2.SideOfLine(a, b), float32(0))

	p3 := F{-0.5, 0.5}
	assert.Greater(t, p3.SideOfLine(a, b), float32(0))
}

func TestAngleBetween(t *testing.T) {
	up := F{0, 1}
	upRight := F{1, 1}
	right := F{1, 0}
	down := F{0, -1}
	downLeft := F{-1, -1}

	assert.Equal(t, float32(math.Pi/2), right.AngleBetween(up))
	assert.Equal(t, float32(math.Pi/4), right.AngleBetween(upRight))
	assert.Equal(t, float32(-math.Pi/2), up.AngleBetween(right))
	assert.Equal(t, float32(math.Pi), up.AngleBetween(down))
	assert.Equal(t, float32(math.Pi*3/4), up.AngleBetween(downLeft))
	assert.Equal(t, float32(-math.Pi*3/4), downLeft.AngleBetween(up))

}
