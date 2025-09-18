package matheval

import (
	"fmt"
)

// A Literal represents a number or a variable
type Literal struct {
	val      float64
	variable string
}

func NewLiteralVal(v float64) *Literal {
	return &Literal{val: v}
}

func NewLiteralVar(v string) *Literal {
	return &Literal{variable: v}
}

func (self *Literal) String() string {
	if self.variable != "" {
		return fmt.Sprintf("%s", self.variable)
	} else {
		return fmt.Sprintf("%.2f", self.val)
	}
}

func (self *Literal) Evaluate(vars map[string]float64) float64 {
	if self.variable == "" {
		return self.val
	} else {
		if val, exists := vars[self.variable]; exists {
			return val
		} else if val, exists := commonConstants[self.variable]; exists {
			return val
		} else {
			return 0
		}
	}

}
