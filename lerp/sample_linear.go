package lerp

import "math"

// SampleLinear lerps between equally spaced values in the [0, 1] range
func SampleLinear(samples []float32, t float32) float32 {
	if t < 0 {
		return samples[0]
	}
	if t > 1 {
		return samples[len(samples)-1]
	}
	step := 1 / float64(len(samples)-1)
	i := math.Floor(float64(t) / step)
	j := math.Ceil(float64(t) / step)
	if i == j {
		return samples[int(i)]
	}
	Xi := float32(i * step)
	Xj := float32(j * step)
	ratio := (t - Xi) / (Xj - Xi)
	return (1-ratio)*samples[int(i)] + ratio*samples[int(j)]
}
