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

const cosSinPrecision = 1024 * 16

var cosSinTable = make([]float32, cosSinPrecision*2)

func init() {
	for i := 0; i < precision; i++ {
		rad := float64(i) / precision * 2 * math.Pi
		cosTable[i] = float32(math.Cos(rad))
	}
	for i := 0; i < cosSinPrecision; i++ {
		rad := float64(i) / cosSinPrecision * 2 * math.Pi
		cosSinTable[2*i] = float32(math.Cos(rad))
		cosSinTable[2*i+1] = -float32(math.Sin(rad))
	}
}

// Cos returns an interpolated value with accuracy <0.0001
func Cos(angle float32) float32 {
	floatIndex, frac := Modf(float32(math.Abs(float64(angle/2/math.Pi))) * precision)
	index := int32(floatIndex)

	low := index % precision
	high := (low + 1) % precision

	return cosTable[low]*(1-frac) + cosTable[high]*(frac)
}

// Sin returns an interpolated value with accuracy <0.0001
func Sin(angle float32) float32 {
	angle = math.Pi/2 - angle
	floatIndex, frac := Modf(float32(math.Abs(float64(angle/2/math.Pi))) * precision)
	index := int32(floatIndex)

	low := index % precision
	high := (low + 1) % precision

	return -cosTable[low]*(1-frac) - cosTable[high]*(frac)
}

// CosSin returns interpolated values with accuracy <0.0001
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

// CosSinFast is superfast, but doesn't interpolate, so it will return the same value for multiple close input angles
// The error is <0.0004
func CosSinFast(angle float32) (cos, sin float32) {
	index := int(angle*cosSinPrecision/2/math.Pi) % cosSinPrecision
	// handle negative angles
	index = (index + cosSinPrecision) % cosSinPrecision

	return cosSinTable[2*index], cosSinTable[2*index+1]
}

// CosSinFastD is superfast, but doesn't interpolate, so it will return the same value for multiple close input angles
// The error is <0.0004
func CosSinFastD(angle float64) (cos, sin float64) {
	index := int(angle*cosSinPrecision/2/math.Pi) % cosSinPrecision
	// handle negative angles
	index = (index + cosSinPrecision) % cosSinPrecision

	return float64(cosSinTable[2*index]), float64(cosSinTable[2*index+1])
}
