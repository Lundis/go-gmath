package matheval

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
)

type builtinFunc struct {
	id     string
	params []*Node
}

var oneParamFuncs = []string{"cos", "sin", "sqrt", "abs"}

var twoParamFuncs = []string{"mod", "pow"}

var anyParamFuncs = []string{"min", "max"}

func newBuiltinFunc(id string, params []*Node) (*builtinFunc, error) {
	expectedParams := 0
	if slices.Index(oneParamFuncs, id) >= 0 {
		expectedParams = 1
	} else if slices.Index(twoParamFuncs, id) >= 0 {
		expectedParams = 2
	} else if slices.Index(anyParamFuncs, id) >= 0 {
		expectedParams = 99
	}
	if expectedParams == 0 {
		return nil, errors.New(fmt.Sprintf("Unknown function: %v", id))
	}

	if expectedParams <= 2 && len(params) != expectedParams {
		return nil, fmt.Errorf("Wrong number of parameters for function %v: %v", id, len(params))
	}
	if expectedParams == 99 && len(params) < 2 {
		return nil, fmt.Errorf("Function %v requires at least two parameters", id)
	}

	f := new(builtinFunc)
	f.id = id
	f.params = params
	return f, nil
}

func (self *builtinFunc) String() string {
	if len(self.params) == 0 {
		return fmt.Sprintf("%v()", self.id)
	} else {
		inside_strs := make([]string, len(self.params))
		for i, n := range self.params {
			inside_strs[i] = n.String()
		}
		args := strings.Join(inside_strs, ", ")
		return fmt.Sprintf("%v(%v)", self.id, args)
	}

}

func (self *builtinFunc) Evaluate(vars map[string]float64) float64 {
	r := 0.0
	params := make([]float64, len(self.params))
	for i, n := range self.params {
		params[i] = n.Evaluate(vars)
	}
	switch self.id {
	case "cos":
		r = math.Cos(params[0])
	case "sin":
		r = math.Sin(params[0])
	case "abs":
		r = math.Abs(params[0])
	case "sqrt":
		r = math.Sqrt(params[0])
	case "mod":
		r = math.Mod(params[0], params[1])
	case "pow":
		r = math.Pow(params[0], params[1])
	case "min":
		r := params[0]
		for _, p := range params[1:] {
			if p < r {
				r = p
			}
		}

	case "max":
		r := params[0]
		for _, p := range params[1:] {
			if p > r {
				r = p
			}
		}
	default:
		panic("not implemented")
	}
	return r
}
