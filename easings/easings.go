package easings

import (
	"math"
)

func EaseOutQuad(x float32) float32 {
	return 1 - (1-x)*(1-x)
}

func EaseInOutQuad(x float32) float32 {
	if x < 0.5 {
		return 2 * x * x
	}
	return 1 - (-2*x+2)*(-2*x+2)/2
}

func EaseOutN(x, factor float32) float32 {
	return 1 - float32(math.Pow(float64(1-x), float64(factor)))
}

func SmoothStep(x float32) float32 {
	return x * x * (3 - 2*x)
}

func EaseOutBounce(x float32) float32 {
	const n1 = 7.5625
	const d1 = 2.75

	if x < 1/d1 {
		return n1 * x * x
	} else if x < 2/d1 {
		x -= 1.5 / d1
		return n1*x*x + 0.75
	} else if x < 2.5/d1 {
		x -= 2.25 / d1
		return n1*x*x + 0.9375
	} else {
		x -= 2.625 / d1
		return n1*x*x + 0.984375
	}
}

func EaseInBounce(x float32) float32 {
	return 1 - EaseOutBounce(1-x)
}
