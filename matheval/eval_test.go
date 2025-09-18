package matheval

import (
	"github.com/Lundis/go-gmath/fastmath"
	"math"
	"testing"
)

func TestEvalCommon(t *testing.T) {
	tests := map[string]float64{
		"1 + 2 + 3":       6,
		"pi + 1":          math.Pi + 1,
		"e + 1":           math.E + 1,
		"e / pi":          math.E / math.Pi,
		"(e + 2*pi) / pi": (math.E + 2*math.Pi) / math.Pi,
		"abs(-5)":         5,
	}
	for str, expected := range tests {
		result, err := Eval(str)
		if err != nil {
			t.Error(err.Error())
		} else if !fastmath.Equald(result, expected, 0.001) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestEvalVariables(t *testing.T) {
	const a = 5
	const b = 23
	const x = 12.34
	vars := map[string]float64{
		"a": a,
		"b": b,
		"x": x,
	}
	tests := map[string]float64{
		"1 + 2 + a": 1 + 2 + a,
		"pi + x/b":  math.Pi + x/b,
		"abs(-a)":   a,
	}
	for str, expected := range tests {
		result, err := EvalVariables(str, vars)
		if err != nil {
			t.Error(err.Error())
		} else if !fastmath.Equald(result, expected, 0.001) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}
