package visitor

import "math"

/*
 访问者模式
*/

//图形形状接口:accept接受访问者【Element类】
type Shape interface {
	Accept(visitor ShapeVisitor)
}

//Circle【ConcreteElement类】
type Circle struct {
	radius float32
}

//圆：实现accept
func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

//Rectangle【ConcreteElement类】
type Rectangle struct {
	width  float32
	height float32
}

//矩形：实现accept
func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

//图形访问者接口【抽象Visitor】
type ShapeVisitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}

//AreaCalculator【Visitor类】
type AreaCalculator struct {
	area float32
}

//计算圆面积
func (calc *AreaCalculator) VisitCircle(circle *Circle) {
	calc.area += math.Pi * circle.radius * circle.radius
}

func (calc *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	calc.area += rectangle.width * rectangle.height
}

func (calc *AreaCalculator) GetArea() float32 {
	return calc.area
}

func NewVisitor() *AreaCalculator {
	return &AreaCalculator{area: 0.00}
}
