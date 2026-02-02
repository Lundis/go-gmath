package vec2

import (
	"math"
	"math/rand/v2"
	"strconv"

	"github.com/Lundis/go-gmath/fastmath"
)

type D struct {
	X, Y float64
}

func NewPolarD(angle, radius float64) D {
	cos, sin := fastmath.CosSinD(angle)
	return D{X: cos, Y: sin}.MulScalar(radius)
}

func NewRandomD(minValue, maxValue float64) D {
	spread := maxValue - minValue
	return D{
		rand.Float64(),
		rand.Float64(),
	}.MulScalar(spread).AddScalar(minValue)
}

func (v D) AsInt() I {
	return I{X: int32(v.X), Y: int32(v.Y)}
}

func (v D) AsFloat() F {
	return F{X: float32(v.X), Y: float32(v.Y)}
}

func (v D) Equals(other D) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v D) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v D) String() string {
	var xString, yString string
	if v.X == float64(int(v.X)) {
		xString = strconv.Itoa(int(v.X))
	} else {
		xString = strconv.FormatFloat(v.X, 'f', 4, 64)
	}
	if v.Y == float64(int(v.Y)) {
		yString = strconv.Itoa(int(v.Y))
	} else {
		yString = strconv.FormatFloat(v.Y, 'f', 4, 64)
	}
	return "(" + xString + ", " + yString + ")"
}

func (v D) Plus(other D) D {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v D) AddScalar(scalar float64) D {
	return v.AddScalars(scalar, scalar)
}

func (v D) AddScalars(x, y float64) D {
	v.X += x
	v.Y += y
	return v
}

func (v D) Minus(other D) D {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v D) SubScalar(scalar float64) D {
	return v.SubScalars(scalar, scalar)
}

func (v D) SubScalars(x, y float64) D {
	v.X -= x
	v.Y -= y
	return v
}

func (v D) Mul(other D) D {
	v.X *= other.X
	v.Y *= other.Y
	return v
}

func (v D) MulScalar(scalar float64) D {
	return v.MulScalars(scalar, scalar)
}

func (v D) MulScalars(x, y float64) D {
	v.X *= x
	v.Y *= y
	return v
}

func (v D) Div(other D) D {
	v.X /= other.X
	v.Y /= other.Y
	return v
}

func (v D) DivScalar(scalar float64) D {
	return v.DivScalars(scalar, scalar)
}

func (v D) DivScalars(x, y float64) D {
	v.X /= x
	v.Y /= y
	return v
}

func (v D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v D) DistanceTo(v2 D) float64 {
	return v.Minus(v2).Magnitude()
}

func (v D) DistanceToLine(a, b D) float64 {
	ab := b.Minus(a)
	ap := v.Minus(a)

	cross := ab.Cross(ap)

	if cross < 0 {
		cross = -cross
	}

	return cross / ab.Magnitude()
}

// SideOfLine calculates which side of the line A->B the point P lies on. Check the sign of the response.
func (v D) SideOfLine(a, b D) float64 {
	ab := b.Minus(a)
	ap := v.Minus(a)

	return ab.Cross(ap)
}

func (v D) DistanceToSquared(v2 D) float64 {
	diff := v.Minus(v2)
	return diff.X*diff.X + diff.Y*diff.Y
}

func (v D) Normalized() D {
	m := v.Magnitude()

	if m > 0.0 {
		return v.DivScalar(m)
	} else {
		return v
	}
}

func (v D) Angle() float64 {
	return fastmath.Atan2D(v.Y, v.X)
}

// AngleBetweenLines calculates the angle between two lines starting at origo
// returns values in the range [-Pi, Pi].
func (v D) AngleBetweenLines(v2 D) float64 {
	angle := v2.Angle() - v.Angle()
	if angle > math.Pi {
		angle -= 2 * math.Pi
	} else if angle <= -math.Pi {
		angle += 2 * math.Pi
	}
	return angle
}

// AngleTo returns the angle of the line v->v2
func (v D) AngleTo(v2 D) float64 {
	return v2.Minus(v).Angle()
}

func (v D) Abs() D {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
	return v
}

func (v D) Clamp(low, high D) D {
	return low.Max(v.Min(high))
}

func (v D) Min(v2 D) D {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
	return v
}
func (v D) Max(v2 D) D {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
	return v
}

func (v D) Round() D {
	v.X = math.Round(v.X)
	v.Y = math.Round(v.Y)
	return v
}

func (v D) Floor() D {
	v.X = float64(int(v.X))
	v.Y = float64(int(v.Y))
	return v
}

func (v D) Ceil() D {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	return v
}

func (v D) Swap() D {
	v.X, v.Y = v.Y, v.X
	return v
}

func (v D) Perpendicular() D {
	v.X, v.Y = v.Y, -v.X
	return v
}

func (v D) WithX(value float64) D {
	v.X = value
	return v
}

func (v D) WithY(value float64) D {
	v.Y = value
	return v
}

func (v D) NegatedY() D {
	v.Y = -v.Y
	return v
}

func (v D) Components() (x, y float64) {
	return v.X, v.Y
}

func (v D) Rotate(angle float64) D {
	cos, sin := fastmath.CosSinD(angle)
	return D{
		X: v.X*cos + v.Y*-sin,
		Y: v.X*sin + v.Y*cos,
	}
}

func (v D) IsBetweenInclusive(left, right D) bool {
	return left.X <= v.X && v.X <= right.X &&
		left.Y <= v.Y && v.Y <= right.Y
}

func (v D) Cross(other D) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v D) Dot(other D) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v D) Reflect(other D) D {
	factor := -2 * v.Dot(other)
	return D{
		X: factor*v.X + other.X,
		Y: factor*v.Y + other.Y,
	}
}
