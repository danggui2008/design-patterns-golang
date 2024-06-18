package strategy

/*
策略模式
*/

//MathOperation【抽象策略类（Strategy）】
type MathOperation interface {
	Operate(a float32, b float32) float32
}

//Addition【具体策略类ConcreteStrategy】
type Addition struct{}

//加法：实现策略接口
func (add *Addition) Operate(a float32, b float32) float32 {
	return a + b
}

//【具体策略类ConcreteStrategy】
type Subtraction struct{}

//MathOperation减法：实现策略接口
func (sub *Subtraction) Operate(a float32, b float32) float32 {
	return a - b
}

//Multiplication【具体策略类ConcreteStrategy】
type Multiplication struct{}

//乘法：实现策略接口
func (mul *Multiplication) Operate(a float32, b float32) float32 {
	return a * b
}

//Calculator【上下文Context】
type Calculator struct {
	operation MathOperation
}

func NewCalculator(operation MathOperation) *Calculator {
	return &Calculator{operation: operation}
}

func (calc *Calculator) SetOperation(operation MathOperation) {
	calc.operation = operation
}

func (calc *Calculator) PerformOperation(a float32, b float32) float32 {
	return calc.operation.Operate(a, b)
}
