package lerp

import (
	"github.com/lundis/go-gmath/vec2"
)

func Lerp2(v1, v2 vec2.F, t float32) vec2.F {
	return vec2.F{
		X: v1.X + (v2.X-v1.X)*t,
		Y: v1.Y + (v2.Y-v1.Y)*t,
	}
}

func Lerp(v1, v2 float32, t float32) float32 {
	return v1 + (v2-v1)*t
}
