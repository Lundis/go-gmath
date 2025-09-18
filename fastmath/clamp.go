package fastmath

func Clamp(x, lower, upper float32) float32 {
	return max(lower, min(upper, x))
}
func ClampI(x, lower, upper int32) int32 {
	return max(lower, min(upper, x))
}
func ClampD(x, lower, upper float64) float64 {
	return max(lower, min(upper, x))
}
