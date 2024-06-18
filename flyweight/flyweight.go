package flyweight

import "fmt"

/*
	flyweight享元模式
*/
//Color颜色【内部状态】
type Color struct {
	color string
}

//Coordinate坐标【外部状态】
type Coordinate struct {
	x float32
	y float32
}

func NewCoordinate(x float32, y float32) Coordinate {
	return Coordinate{x: x, y: y}
}

//Shape【抽象享元角色（Flyweight）】
type Shape interface {
	Draw(coordinate Coordinate)
}

/*
Circle【具体享元角色（Concrete Flyweight）】
圆：画一个圆包括：坐标，颜色
这里我们假设颜色只有：红、绿、蓝，黄，而坐标不固定。
那么圆的坐标适合作为对象的外部状态，而颜色为：内部状态。
*/
type Circle struct {
	color Color
}

func (c *Circle) Draw(coordinate Coordinate) {
	fmt.Printf("Draw a %s color circle at x:%0.2f,y:%0.2f\n", c.color.color, coordinate.x, coordinate.y)
}

type ShapeFactory struct {
	shapes map[string]Shape
}

func NewShapeFactory() ShapeFactory {
	return ShapeFactory{map[string]Shape{}}
}

func (sf *ShapeFactory) GetShape(color string) Shape {
	shape, ok := sf.shapes[color]
	if !ok {
		shape = &Circle{Color{color}}
		sf.shapes[color] = shape
	}
	return shape
}
