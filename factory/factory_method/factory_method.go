package factorymethod

/*
工厂方法模式
*/

import "fmt"

//图形接口【抽象产品（Abstract Product）】
type Shape interface {
	Draw()
}

//圆【具体产品（Concrete Product）】
type Circle struct{}

//Circle实现抽象产品
func (c *Circle) Draw() {
	fmt.Println("Drawing a circle")
}

//矩形【具体产品（Concrete Product）】
type Rectangle struct{}

//Rectangle实现抽象产品
func (rec *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle")
}

//图形工厂抽象接口【抽象工厂（Abstract Factory）】
type ShapeFactory interface {
	CreateShape() Shape
}

//圆工厂：负责创建圆形产品【具体工厂（Concrete Factory）】
type CircleFactory struct{}

func (cf *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

//矩形工厂：负责创建矩形产品【具体工厂（Concrete Factory）】
type RectangleFactory struct{}

func (rf *RectangleFactory) CreateShape() Shape {
	return &Rectangle{}
}
