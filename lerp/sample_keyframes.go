package lerp

import "github.com/Lundis/go-gmath/vec2"

// SampleKeyframes samples by keyframes. X is the time, and Y is the value. keyFrames must be sorted by time
func SampleKeyframes(keyFrames []vec2.F, t float32) float32 {
	for i := 0; i < len(keyFrames); i++ {
		if t > keyFrames[i].X {
			continue
		}
		f1 := keyFrames[i-1]
		f2 := keyFrames[i]

		ratio := (t - f1.X) / (f2.X - f1.X)
		return f1.Y + ratio*(f2.Y-f1.Y)
	}
	// beyond the range
	return keyFrames[len(keyFrames)-1].Y
}
