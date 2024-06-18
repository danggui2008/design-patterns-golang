package template

import "fmt"

/*
	模板方法
*/
type TemplateMethod interface {
	TemplateMethod()
}
type AbstractClass struct {
	impl *ConcreteClass
}

func NewAbstractClass(impl *ConcreteClass) *AbstractClass {
	return &AbstractClass{impl: impl}
}

func (a AbstractClass) TemplateMethod() {
	a.step1()
	a.impl.Step2()
	a.impl.Step3()
	a.step4()

}
func (a AbstractClass) step1() {
	fmt.Println("AbstractClass Step1")
}
func (a AbstractClass) Step2() {
	fmt.Println("AbstractClass Step1")
}
func (a AbstractClass) Step3() {
	fmt.Println("AbstractClass Step1")
}
func (a AbstractClass) step4() {
	fmt.Println("AbstractClass Step1")
}

type ConcreteClass struct {
}

func NewConcreteClass() *ConcreteClass {
	return &ConcreteClass{}
}

func (c ConcreteClass) Step2() {
	fmt.Println("ConcreteClass Step2")
}
func (c ConcreteClass) Step3() {
	fmt.Println("ConcreteClass Step3")
}
