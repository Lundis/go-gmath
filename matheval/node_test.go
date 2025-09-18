package matheval

import (
	"github.com/Lundis/go-gmath/fastmath"
	"math"
	"testing"
)

func TestEvaluateSimple(t *testing.T) {
	assertEquals(t, "1+3", 4, nil)
	assertEquals(t, "1-3", -2, nil)
	assertEquals(t, "2*3", 6, nil)
	assertEquals(t, "2/4", 2.0/4, nil)
}

func TestEvaluatePrecedence(t *testing.T) {
	assertEquals(t, "5 * (3 + 2)", 25, nil)
	assertEquals(t, "5 * 3 + 5 * 6", 45, nil)
	assertEquals(t, "5 * 3 * (3 + 2)", 75, nil)
	assertEquals(t, "2 + 3 * 4", 14, nil)
}

func TestEvaluateVariables(t *testing.T) {
	vars := map[string]float64{
		"a": 5,
		"b": 23,
	}
	assertEquals(t, "a+b", 28, vars)
	assertEquals(t, "a-b", -18, vars)
	assertEquals(t, "a*b", 5*23, vars)
	assertEquals(t, "a/b", 5.0/23, vars)
}

func assertEquals(t *testing.T, expr string, correct float64, vars map[string]float64) {
	node, _ := Parse(expr)
	result := node.Evaluate(vars)
	if !fastmath.Equald(correct, result, 0.0001) {
		t.Errorf("Evaluated %v. Got %v. Expected %v", expr, result, correct)
	}
}

func TestFunction(t *testing.T) {
	// test cases at x=0, x=0.5, x=1
	fun_str := []string{"1 + x/2", "5*x", "pi - pi", "pi + e"}
	expected := [][]float64{{1, 1.25, 1.5}, {0, 2.5, 5}, {0, 0, 0}, {math.Pi + math.E, math.Pi + math.E, math.Pi + math.E}}
	for i := range fun_str {
		node, err := Parse(fun_str[i])
		if err != nil {
			t.Errorf("Failed to parse %v: %v", fun_str[i], err)
		}
		exp := expected[i]
		for x, v := range exp {
			xx := float64(x) / 2
			result := node.Evaluate(map[string]float64{"x": xx})
			if !fastmath.Equald(result, v, 0.0001) {
				t.Errorf("%v, x=%v, expected %v, got %v", fun_str[i], xx, v, result)
			}
		}
	}
}
