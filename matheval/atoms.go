package matheval

type Function func(float64) float64

type Atom interface {
	String() string
	Evaluate(vars map[string]float64) float64
}
