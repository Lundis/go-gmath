package vec2

import "github.com/Lundis/go-gmath/fastmath"

type RotationKernel struct {
	cos, sin float32
}

func NewRotationKernel(angle float32) (rk RotationKernel) {
	rk.cos, rk.sin = fastmath.CosSin(angle)
	return
}

func (rk RotationKernel) Rotate(v F) F {
	return F{
		X: v.X*rk.cos + v.Y*-rk.sin,
		Y: v.X*rk.sin + v.Y*rk.cos,
	}
}
