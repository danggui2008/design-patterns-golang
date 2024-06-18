package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	//3
	numberExpression := NewNumberExpression(3)
	//(10-3)
	subExpression := NewSubExpression(NewNumberExpression(10), NewNumberExpression(3))
	//3+(10-3)
	addExpression := NewAddExpression(numberExpression, subExpression)
	//10
	value := addExpression.Interpret()
	fmt.Printf("value:%f\n", value)

}
