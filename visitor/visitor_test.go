package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	circle := Circle{10.00}
	rectangle := Rectangle{width: 10.00, height: 12.00}

	areaCalculator := NewVisitor()
	circle.Accept(areaCalculator)
	area := areaCalculator.GetArea()
	fmt.Printf("circle area:%f\n", area)

	rectangle.Accept(areaCalculator)
	area = areaCalculator.GetArea()
	fmt.Printf("total area:%f\n", area)
}
