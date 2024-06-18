package bridge

import "fmt"

/*
桥接模式
*/

//颜色接口【实现（Implementor）】
type Color interface {
	ApplyColor()
}

//Red【具体实现（ConcreteImplementor）】
type Red struct{}

func (r *Red) ApplyColor() {
	fmt.Println("applying red color")
}

//Blue【具体实现（ConcreteImplementor）】
type Blue struct{}

func (b *Blue) ApplyColor() {
	fmt.Println("applying blue color")
}

//shape接口【抽象（Abstraction）】
type Shape interface {
	Draw()
}

//Circle【修正抽象（RefinedAbstraction）】
type Circle struct {
	color Color
}

func (c *Circle) Draw() {
	//画一个圆，然后给圆着色
	fmt.Println("Draw a circle")
	c.color.ApplyColor()
}

//Square【修正抽象（RefinedAbstraction）】
type Square struct {
	color Color
}

func (c *Square) Draw() {
	//画一个正方形，然后给圆着色
	fmt.Println("Draw a square")
	c.color.ApplyColor()
}
