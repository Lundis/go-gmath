package vec3

import (
	"math"
	"math/rand/v2"
	"strconv"

	"github.com/Lundis/go-gmath/fastmath"
)

type F struct {
	X, Y, Z float32
}

func NewRandomF(minValue, maxValue float32) F {
	spread := maxValue - minValue
	return F{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}.MulScalar(spread).AddScalar(minValue)
}

func (v F) Equals(other F) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

func (v F) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

func (v F) String() string {
	var xString, yString, zString string
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
	if v.Z == float32(int(v.Z)) {
		zString = strconv.Itoa(int(v.Z))
	} else {
		zString = strconv.FormatFloat(float64(v.Z), 'f', 4, 32)
	}
	return "(" + xString + ", " + yString + ", " + zString + ")"
}

func (v F) Plus(other F) F {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
	return v
}

func (v F) AddScalar(scalar float32) F {
	return v.AddScalars(scalar, scalar, scalar)
}

func (v F) AddScalars(x, y, z float32) F {
	v.X += x
	v.Y += y
	v.Z += z
	return v
}

func (v F) Minus(other F) F {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
	return v
}

func (v F) SubScalar(scalar float32) F {
	return v.SubScalars(scalar, scalar, scalar)
}

func (v F) SubScalars(x, y, z float32) F {
	v.X -= x
	v.Y -= y
	v.Z -= z
	return v
}

func (v F) Mul(other F) F {
	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z
	return v
}

func (v F) MulScalar(scalar float32) F {
	return v.MulScalars(scalar, scalar, scalar)
}

func (v F) MulScalars(x, y, z float32) F {
	v.X *= x
	v.Y *= y
	v.Z *= z
	return v
}

func (v F) Div(other F) F {
	v.X /= other.X
	v.Y /= other.Y
	v.Z /= other.Z
	return v
}

func (v F) DivScalar(scalar float32) F {
	return v.DivScalars(scalar, scalar, scalar)
}

func (v F) DivScalars(x, y, z float32) F {
	v.X /= x
	v.Y /= y
	v.Z /= z
	return v
}

func (v F) Magnitude() float32 {
	return fastmath.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v F) DistanceTo(v2 F) float32 {
	return v.Minus(v2).Magnitude()
}

func (v F) DistanceToLine(a, b F) float32 {
	ab := b.Minus(a)
	ap := v.Minus(a)

	cross := ab.Cross(ap)

	return cross.Magnitude() / ab.Magnitude()
}

func (v F) DistanceToSquared(v2 F) float32 {
	diff := v.Minus(v2)
	return diff.X*diff.X + diff.Y*diff.Y + diff.Z*diff.Z
}

func (v F) Normalized() F {
	m := v.Magnitude()

	if m > 0.0 {
		return v.DivScalar(m)
	} else {
		return v
	}
}

func (v F) Abs() F {
	v.X = float32(math.Abs(float64(v.X)))
	v.Y = float32(math.Abs(float64(v.Y)))
	v.Z = float32(math.Abs(float64(v.Z)))
	return v
}

func (v F) Clamp(low, high F) F {
	return low.Max(v.Min(high))
}

func (v F) Min(v2 F) F {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
	v.Z = min(v.Z, v2.Z)
	return v
}
func (v F) Max(v2 F) F {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
	v.Z = max(v.Z, v2.Z)
	return v
}

func (v F) Round() F {
	v.X = float32(math.Round(float64(v.X)))
	v.Y = float32(math.Round(float64(v.Y)))
	v.Z = float32(math.Round(float64(v.Z)))
	return v
}

func (v F) Floor() F {
	v.X = float32(int(v.X))
	v.Y = float32(int(v.Y))
	v.Z = float32(int(v.Z))
	return v
}

func (v F) Ceil() F {
	v.X = float32(math.Ceil(float64(v.X)))
	v.Y = float32(math.Ceil(float64(v.Y)))
	v.Z = float32(math.Ceil(float64(v.Z)))
	return v
}

func (v F) Components() (x, y, z float32) {
	return v.X, v.Y, v.Z
}

func (v F) Array() [3]float32 {
	return [3]float32{v.X, v.Y, v.Z}
}

func (v F) Rotate(angle float32) F {
	cos, sin := fastmath.CosSin(angle)
	return F{
		X: v.X*cos + v.Y*-sin,
		Y: v.X*sin + v.Y*cos,
	}
}

func (v F) IsBetweenInclusive(left, right F) bool {
	return left.X <= v.X && v.X <= right.X &&
		left.Y <= v.Y && v.Y <= right.Y
}

func (v F) Cross(other F) F {
	return F{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v F) Dot(other F) float32 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}
