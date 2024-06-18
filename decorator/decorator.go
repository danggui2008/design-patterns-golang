package decorator

import "fmt"

/*
装饰器模式
*/

//咖啡接口【组件（Component）】
type Coffee interface {
	Cost() float64
	Desc() string
}

//普通咖啡【具体组件（Concrete Component）】
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
	return 10.00
}

func (s *SimpleCoffee) Desc() string {
	return "Simple Coffee"
}

//MilkDecorator【具体装饰器（Concrete Decorator）】
type MilkDecorator struct {
	Coffee
}

func NewMilkDecorator(coffee Coffee) Coffee {
	return &MilkDecorator{Coffee: coffee}
}

//实现咖啡接口
func (milk *MilkDecorator) Cost() float64 {
	return milk.Coffee.Cost() + 2
}

func (milk *MilkDecorator) Desc() string {
	return fmt.Sprintf("%s,with  Milk", milk.Coffee.Desc())
}

type SugarDecorator struct {
	Coffee
}

func NewSugarDecorator(coffee Coffee) Coffee {
	return &SugarDecorator{Coffee: coffee}
}

//实现咖啡接口
func (sugar *SugarDecorator) Cost() float64 {
	return sugar.Coffee.Cost() + 2
}

func (sugar *SugarDecorator) Desc() string {
	return fmt.Sprintf("%s,with  Sugar", sugar.Coffee.Desc())
}
