package simplefactory

/*
简单工厂模式
*/

import (
	"errors"
	"fmt"
)

// 图形接口【抽象产品（Abstract Product）】
type Shape interface {
	Draw()
}

// 圆【具体产品（Concrete Product）】
type Circle struct{}

// 实现产品接口
func (c *Circle) Draw() {
	fmt.Println("Drawing a circle")
}

// 矩形【具体产品（Concrete Product）】
type Rectangle struct{}

// 实现产品接口
func (rec *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle")
}

// 工厂【简单工厂】
func NewShape(name string) (Shape, error) {
	if name == "Circle" {
		return &Circle{}, nil
	} else if name == "Rectangle" {
		return &Rectangle{}, nil
	} else {
		return nil, errors.New("create shape error")
	}
}
