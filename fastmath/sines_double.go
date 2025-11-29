package fastmath

import (
	"math"
)

var cosTableD = make([]float64, precision)

func init() {
	for i := 0; i < precision; i++ {
		rad := float64(i) / precision * 2 * math.Pi
		cosTableD[i] = math.Cos(rad)
	}
}

func CosD(angle float64) float64 {
	floatIndex, frac := math.Modf(math.Abs(angle/2/math.Pi) * precision)
	index := int32(floatIndex)

	low := index % precision
	high := (low + 1) % precision

	return cosTableD[low]*(1-frac) + cosTableD[high]*(frac)
}

func SinD(angle float64) float64 {
	angle = math.Pi/2 - angle
	floatIndex, frac := math.Modf(math.Abs(angle/2/math.Pi) * precision)
	index := int32(floatIndex)

	low := index % precision
	high := (low + 1) % precision

	return -cosTableD[low]*(1-frac) - cosTableD[high]*(frac)
}

func CosSinD(angle float64) (cos, sin float64) {
	floatIndex, frac := math.Modf(math.Abs(angle/2/math.Pi) * precision)
	index := int32(floatIndex)

	low := index % precision
	high := low + 1
	if high >= precision {
		high -= precision
	}
	var lowSin = -low
	if angle < 0 {
		lowSin += precision - precision/4
	} else {
		lowSin += precision / 4
	}
	if lowSin < 0 {
		lowSin += precision
	}
	highSin := lowSin - 1
	if highSin < 0 {
		highSin += precision
	}

	return cosTableD[low]*(1-frac) + cosTableD[high]*frac,
		-cosTableD[lowSin]*(1-frac) - cosTableD[highSin]*(frac)
}
