package matheval

import (
	"math"
	"strings"
)

var commonConstants = map[string]float64{
	"pi": math.Pi,
	"e":  math.E,
}

func Eval(expr string) (float64, error) {
	expr = strings.ToLower(expr)
	return EvalVariables(expr, commonConstants)
}

func EvalVariables(expr string, vars map[string]float64) (float64, error) {
	node, err := Parse(expr)
	if err != nil {
		return 0, err
	} else {
		return node.Evaluate(vars), nil
	}
}
