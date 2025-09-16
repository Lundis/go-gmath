package vec2

import (
	"github.com/lundis/go-gmath/fastmath"
	"math"
	"math/rand/v2"
	"strconv"
)

type F struct {
	X, Y float32
}

func NewPolarF(angle, radius float32) F {
	cos, sin := fastmath.CosSin(angle)
	return F{X: cos, Y: sin}.MulScalar(radius)
}

func NewRandomF(minValue, maxValue float32) F {
	spread := maxValue - minValue
	return F{
		rand.Float32(),
		rand.Float32(),
	}.MulScalar(spread).AddScalar(minValue)
}

func (v F) AsInt() I {
	return I{X: int32(v.X), Y: int32(v.Y)}
}

func (v F) Equals(other F) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v F) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v F) String() string {
	var xString, yString string
	if v.X == float32(int(v.X)) {
		xString = strconv.Itoa(int(v.X))
	} else {
		xString = strconv.FormatFloat(float64(v.X), 'f', 4, 32)
	}
	if v.Y == float32(int(v.Y)) {
		yString = strconv.Itoa(int(v.Y))
	} else {
		yString = strconv.FormatFloat(float64(v.Y), 'f', 4, 32)
	}
	return "(" + xString + ", " + yString + ")"
}

// Self is useful if you want to build complex stuff with interfaces
func (v F) Self() F {
	return v
}

func (v F) Plus(other F) F {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v F) AddScalar(scalar float32) F {
	return v.AddScalars(scalar, scalar)
}

func (v F) AddScalars(x, y float32) F {
	v.X += x
	v.Y += y
	return v
}

func (v F) Minus(other F) F {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v F) SubScalar(scalar float32) F {
	return v.SubScalars(scalar, scalar)
}

func (v F) SubScalars(x, y float32) F {
	v.X -= x
	v.Y -= y
	return v
}

func (v F) Mul(other F) F {
	v.X *= other.X
	v.Y *= other.Y
	return v
}

func (v F) MulScalar(scalar float32) F {
	return v.MulScalars(scalar, scalar)
}

func (v F) MulScalars(x, y float32) F {
	v.X *= x
	v.Y *= y
	return v
}

func (v F) Div(other F) F {
	v.X /= other.X
	v.Y /= other.Y
	return v
}

func (v F) DivScalar(scalar float32) F {
	return v.DivScalars(scalar, scalar)
}

func (v F) DivScalars(x, y float32) F {
	v.X /= x
	v.Y /= y
	return v
}

func (v F) Magnitude() float32 {
	return fastmath.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v F) DistanceTo(v2 F) float32 {
	return v.Minus(v2).Magnitude()
}

func (v F) DistanceToLine(a, b F) float32 {
	ab := b.Minus(a)
	ap := v.Minus(a)

	cross := ab.Cross(ap)

	if cross < 0 {
		cross = -cross
	}

	return cross / ab.Magnitude()
}

// SideOfLine calculates which side of the line A->B the point P lies on. Check the sign of the response.
func (v F) SideOfLine(a, b F) float32 {
	ab := b.Minus(a)
	ap := v.Minus(a)

	return ab.Cross(ap)
}

func (v F) DistanceToSquared(v2 F) float32 {
	diff := v.Minus(v2)
	return diff.X*diff.X + diff.Y*diff.Y
}

func (v F) Normalized() F {
	m := v.Magnitude()

	if m > 0.0 {
		return v.DivScalar(m)
	} else {
		return v
	}
}

func (v F) Angle() float32 {
	return fastmath.Atan2(v.Y, v.X)
}

// AngleBetween calculates the angle between two vectors, returning a value in the range [-Pi, Pi].
func (v F) AngleBetween(v2 F) float32 {
	angle := v.Angle() - v2.Angle()
	if angle > math.Pi {
		angle -= 2 * math.Pi
	} else if angle <= -math.Pi {
		angle += 2 * math.Pi
	}
	return angle
}

func (v F) Abs() F {
	v.X = float32(math.Abs(float64(v.X)))
	v.Y = float32(math.Abs(float64(v.Y)))
	return v
}

func (v F) Clamp(low, high F) F {
	return low.Max(v.Min(high))
}

func (v F) Min(v2 F) F {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
	return v
}
func (v F) Max(v2 F) F {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
	return v
}

func (v F) Round() F {
	v.X = float32(math.Round(float64(v.X)))
	v.Y = float32(math.Round(float64(v.Y)))
	return v
}

func (v F) Floor() F {
	v.X = float32(int(v.X))
	v.Y = float32(int(v.Y))
	return v
}

func (v F) Ceil() F {
	v.X = float32(math.Ceil(float64(v.X)))
	v.Y = float32(math.Ceil(float64(v.Y)))
	return v
}

func (v F) Swap() F {
	v.X, v.Y = v.Y, v.X
	return v
}

func (v F) Perpendicular() F {
	v.X, v.Y = v.Y, -v.X
	return v
}

func (v F) WithX(value float32) F {
	v.X = value
	return v
}

func (v F) WithY(value float32) F {
	v.Y = value
	return v
}

func (v F) Components() (x, y float32) {
	return v.X, v.Y
}

func (v F) Rotate(angle float32) F {
	cos, sin := fastmath.CosSin(angle)
	return F{
		X: v.X*cos + v.Y*-sin,
		Y: v.X*sin + v.Y*cos,
	}
}

func (v F) IsBetween(left, right F) bool {
	return left.X <= v.X && v.X <= right.X &&
		left.Y <= v.Y && v.Y <= right.Y
}

func (v F) Cross(other F) float32 {
	return v.X*other.Y - v.Y*other.X
}

func (v F) Dot(other F) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v F) Reflect(other F) F {
	factor := -2 * v.Dot(other)
	return F{
		X: factor*v.X + other.X,
		Y: factor*v.Y + other.Y,
	}
}
