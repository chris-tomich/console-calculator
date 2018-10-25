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
	A Operation
	B Operation
}

func (op *Addition) Evaluate() int {
	return op.A.Evaluate() + op.B.Evaluate()
}

var _ InfixOperation = &Subtraction{}

type Subtraction struct {
	A Operation
	B Operation
}

func (op *Subtraction) Evaluate() int {
	return op.A.Evaluate() - op.B.Evaluate()
}

var _ InfixOperation = &Multiplication{}

type Multiplication struct {
	A Operation
	B Operation
}

func (op *Multiplication) Evaluate() int {
	return op.A.Evaluate() * op.B.Evaluate()
}

var _ InfixOperation = &Division{}

type Division struct {
	A Operation
	B Operation
}

func (op *Division) Evaluate() int {
	return op.A.Evaluate() / op.B.Evaluate()
}
