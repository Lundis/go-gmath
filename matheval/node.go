package matheval

import (
	"fmt"
	"math"
	"strings"
)

type Operator int

const (
	PLUS Operator = iota
	MINUS
	MULT
	DIV
	ATOM
	FUNC
)

type Node struct {
	op    Operator
	nodes []*Node
	data  Atom
}

func (self *Node) childrenStrings() []string {
	strs := make([]string, len(self.nodes))
	for i := range strs {
		strs[i] = self.nodes[i].String()
	}
	return strs
}

func (self *Node) String() string {
	switch self.op {
	case PLUS:
		return "(" + strings.Join(self.childrenStrings(), " + ") + ")"
	case MINUS:
		return " - (" + self.nodes[0].String() + ")"
	case MULT:
		return strings.Join(self.childrenStrings(), " * ")
	case DIV:
		return self.nodes[0].String() + " / " + self.nodes[1].String()
	case ATOM:
		return self.data.String()
	case FUNC:
		return self.data.String() + strings.Join(self.childrenStrings(), ", ") + ")"
	default:
		panic(fmt.Sprintf("Unknown NodeImpl type in String(): %v", self.op))
	}
}

func NewPlusNode(nodes []*Node) *Node {
	n := new(Node)
	n.op = PLUS
	n.nodes = nodes
	return n
}

// Wraps n in a minus node
func NewMinusNode(n *Node) *Node {
	m := new(Node)
	m.op = MINUS
	m.nodes = make([]*Node, 1)
	m.nodes[0] = n
	return m
}

func NewMultNode(nodes []*Node) *Node {
	n := new(Node)
	n.op = MULT
	n.nodes = nodes
	return n
}

func NewDivNode(n1, n2 *Node) *Node {
	m := new(Node)
	m.op = DIV
	m.nodes = make([]*Node, 2)
	m.nodes[0] = n1
	m.nodes[1] = n2
	return m
}

func NewFunctionNode(a Atom) *Node {
	n := new(Node)
	n.op = FUNC
	n.data = a
	return n
}

func NewLiteralNode(num float64) *Node {
	m := new(Node)
	m.op = ATOM
	m.data = NewLiteralVal(num)
	return m
}

func NewVarNode(id string) *Node {
	m := new(Node)
	m.op = ATOM
	m.data = NewLiteralVar(id)
	return m
}

func (self *Node) Evaluate(vars map[string]float64) float64 {
	switch self.op {
	case PLUS:
		sum := 0.0
		for _, v := range self.nodes {
			sum += v.Evaluate(vars)
		}
		return sum
	case MINUS:
		return -self.nodes[0].Evaluate(vars)
	case MULT:
		prod := 1.0
		for _, v := range self.nodes {
			prod *= v.Evaluate(vars)
		}
		return prod
	case DIV:
		first := self.nodes[0].Evaluate(vars)
		second := self.nodes[1].Evaluate(vars)
		if second == 0 {
			return math.Copysign(math.Inf(1), first)
		}
		if math.IsInf(second, 0) {
			return 0
		}
		return first / second
	case ATOM:
		fallthrough
	case FUNC:
		return self.data.Evaluate(vars)
	default:
		return 0
	}
}
