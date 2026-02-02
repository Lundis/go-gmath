package vec2

import (
	"fmt"
	"math"
)

type I struct {
	X, Y int32
}

func (i I) AsFloat() F {
	return F{X: float32(i.X), Y: float32(i.Y)}
}

func (i I) AsDouble() D {
	return D{X: float64(i.X), Y: float64(i.Y)}
}

func (i I) IsZero() bool {
	return i.X == 0 && i.Y == 0
}

func (i I) Components() (x, y int32) {
	return i.X, i.Y
}

func (i I) Add(other I) I {
	return I{X: i.X + other.X, Y: i.Y + other.Y}
}

func (i I) AddScalars(x, y int32) I {
	return I{X: i.X + x, Y: i.Y + y}
}

func (i I) Sub(other I) I {
	return I{X: i.X - other.X, Y: i.Y - other.Y}
}

func (i I) Magnitude() float64 {
	return math.Sqrt(float64(i.X*i.X + i.Y*i.Y))
}

func (i I) Area() int32 {
	return i.X * i.Y
}

func (i I) Equals(other I) bool {
	return i.X == other.X && i.Y == other.Y
}

func (i I) String() string {
	return fmt.Sprintf("(%d, %d)", i.X, i.Y)
}

func (i I) IsBetweenInclusive(left, right I) bool {
	return left.X <= i.X && i.X <= right.X &&
		left.Y <= i.Y && i.Y <= right.Y
}

func (i I) Clamp(left, right I) I {
	if i.X < left.X {
		i.X = left.X
	}
	if i.X > right.X {
		i.X = right.X
	}
	if i.Y < left.Y {
		i.Y = left.Y
	}
	if i.Y > right.Y {
		i.Y = right.Y
	}

	return i
}

func (i I) MinMax(other I) (min_, max_ I) {
	if i.X > other.X {
		max_.X = i.X
		min_.X = other.X
	} else {
		max_.X = other.X
		min_.X = i.X
	}
	if i.Y > other.Y {
		max_.Y = i.Y
		min_.Y = other.Y
	} else {
		max_.Y = other.Y
		min_.Y = i.Y
	}

	return
}

func (i I) Min(other I) I {
	if i.X > other.X {
		i.X = other.X
	}
	if i.Y > other.Y {
		i.Y = other.Y
	}
	return i
}

func (i I) Max(other I) I {
	if i.X < other.X {
		i.X = other.X
	}
	if i.Y < other.Y {
		i.Y = other.Y
	}
	return i
}

func (i I) Abs() I {
	if i.X < 0 {
		i.X = -i.X
	}
	if i.Y < 0 {
		i.Y = -i.Y
	}
	return i
}

func (i I) Index(width int32) int32 {
	return width*i.Y + i.X
}
