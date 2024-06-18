package interpreter

/*
	解释器模式
*/
//Expression【抽象表达式（Abstract Expression）】
type Expression interface {
	Interpret() float32
}

//NumberExpression【终结符表达式（Terminal Expression）】
type NumberExpression struct {
	value float32
}

func NewNumberExpression(value float32) Expression {
	return &NumberExpression{value: value}
}

//实现表达式接口：NumberExpression
func (n *NumberExpression) Interpret() float32 {
	return n.value
}

//加法表达式AddExpression【非终结符表达式（Non-terminal Expression）】
type AddExpression struct {
	leftOperand  Expression
	rightOperand Expression
}

func NewAddExpression(leftOperand Expression, rightOperand Expression) Expression {
	return &AddExpression{leftOperand: leftOperand, rightOperand: rightOperand}
}

func (a *AddExpression) Interpret() float32 {
	return a.leftOperand.Interpret() + a.rightOperand.Interpret()
}

//SubExpression减法表达式【非终结符表达式（Non-terminal Expression）】
type SubExpression struct {
	leftOperand  Expression
	rightOperand Expression
}

func NewSubExpression(leftOperand Expression, rightOperand Expression) Expression {
	return &SubExpression{leftOperand: leftOperand, rightOperand: rightOperand}
}

func (a *SubExpression) Interpret() float32 {
	return a.leftOperand.Interpret() - a.rightOperand.Interpret()
}
