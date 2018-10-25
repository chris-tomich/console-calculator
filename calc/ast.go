package calc

type Operation interface {
	Evaluate() int
}

var _ Operation = &Constant{}

type Constant struct {
	A int
}

func (op *Constant) Evaluate() int {
	return op.A
}

type InfixOperation interface {
	Evaluate() int
}

var _ InfixOperation = &Addition{}

type Addition struct {
	A int
	B int
}

func (op *Addition) Evaluate() int {
	return op.A + op.B
}

var _ InfixOperation = &Subtraction{}

type Subtraction struct {
	A int
	B int
}

func (op *Subtraction) Evaluate() int {
	return op.A - op.B
}

var _ InfixOperation = &Multiplication{}

type Multiplication struct {
	A int
	B int
}

func (op *Multiplication) Evaluate() int {
	return op.A * op.B
}

var _ InfixOperation = &Division{}

type Division struct {
	A int
	B int
}

func (op *Division) Evaluate() int {
	return op.A / op.B
}
