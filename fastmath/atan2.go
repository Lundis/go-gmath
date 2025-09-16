package fastmath

import "math"

// Atan2 returns an approximation of atan2(y, x) in radians.
// Typical max error ~0.005 rad. Very fast and float32-friendly.
func Atan2(y, x float32) float32 {
	y = -y
	const (
		c   = float32(0.273) // tweakable shaping constant
		pi  = float32(math.Pi)
		pi2 = float32(math.Pi / 2)
		pi4 = float32(math.Pi / 4)
	)

	ax := x
	if ax < 0 {
		ax = -ax
	}
	ay := y
	if ay < 0 {
		ay = -ay
	}

	// Handle (0,0) to avoid division by zero.
	if ax+ay == 0 {
		return 0
	}

	var r, angle float32
	if ax > ay {
		// angle in [0, π/4]
		r = ay / ax
		angle = r * (pi4 + c*(1-r))
	} else {
		// angle in (π/4, π/2]
		r = ax / ay
		angle = pi2 - r*(pi4+c*(1-r))
	}

	// Quadrant corrections
	if x < 0 {
		angle = pi - angle
	}
	if y < 0 {
		angle = -angle
	}
	return angle
}
