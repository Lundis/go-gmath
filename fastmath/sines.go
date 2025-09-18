package fastmath

import (
	"math"
)

// SafeAngle64 converts a potentially huge angle value to one that safely fits in a float32
func SafeAngle64(angle float64) float32 {
	div := float64(int64(angle / (2 * math.Pi)))
	return float32(angle - div*2*math.Pi)
}

const precision = 256

var cosTable = make([]float32, precision)

func init() {
	for i := 0; i < precision; i++ {
		rad := float64(i) / precision * 2 * math.Pi
		cosTable[i] = float32(math.Cos(rad))
	}
}

func Cos(angle float32) float32 {
	floatIndex, frac := Modf(float32(math.Abs(float64(angle/2/math.Pi))) * precision)
	index := int32(floatIndex)

	low := index % precision
	high := (low + 1) % precision

	return cosTable[low]*(1-frac) + cosTable[high]*(frac)
}

func Sin(angle float32) float32 {
	angle = math.Pi/2 - angle
	floatIndex, frac := Modf(float32(math.Abs(float64(angle/2/math.Pi))) * precision)
	index := int32(floatIndex)

	low := index % precision
	high := (low + 1) % precision

	return -cosTable[low]*(1-frac) - cosTable[high]*(frac)
}

func CosSin(angle float32) (cos, sin float32) {
	floatIndex, frac := Modf(float32(math.Abs(float64(angle/2/math.Pi))) * precision)
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

	return cosTable[low]*(1-frac) + cosTable[high]*frac,
		-cosTable[lowSin]*(1-frac) - cosTable[highSin]*(frac)
}
