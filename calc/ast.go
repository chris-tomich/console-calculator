package calc

import (
	"strconv"
)

type Operation interface {
	Evaluate() int
	Value() string
}

var _ Operation = &Constant{}

type Constant struct {
	A int
}

func (op *Constant) Evaluate() int {
	return op.A
}

func (op *Constant) Value() string {
	return strconv.Itoa(op.A)
}

type InfixOperation interface {
	Evaluate() int
	Value() string
}

var _ InfixOperation = &Addition{}

type Addition struct {
	A Operation
	B Operation
}

func (op *Addition) Evaluate() int {
	return op.A.Evaluate() + op.B.Evaluate()
}

func (op *Addition) Value() string {
	return "+"
}

var _ InfixOperation = &Subtraction{}

type Subtraction struct {
	A Operation
	B Operation
}

func (op *Subtraction) Evaluate() int {
	return op.A.Evaluate() - op.B.Evaluate()
}

func (op *Subtraction) Value() string {
	return "-"
}

var _ InfixOperation = &Multiplication{}

type Multiplication struct {
	A Operation
	B Operation
}

func (op *Multiplication) Evaluate() int {
	return op.A.Evaluate() * op.B.Evaluate()
}

func (op *Multiplication) Value() string {
	return "*"
}

var _ InfixOperation = &Division{}

type Division struct {
	A Operation
	B Operation
}

func (op *Division) Evaluate() int {
	return op.A.Evaluate() / op.B.Evaluate()
}

func (op *Division) Value() string {
	return "/"
}
