package template

import "testing"

func TestTemplate(t *testing.T) {
	concreteClass := NewConcreteClass()
	abstractClass := NewAbstractClass(concreteClass)
	abstractClass.TemplateMethod()
}
