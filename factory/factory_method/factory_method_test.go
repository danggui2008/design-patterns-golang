package factorymethod

import "testing"

func TestCircle(t *testing.T) {
	factory := CircleFactory{}
	shape := factory.CreateShape()
	if shape != nil {
		shape.Draw()
	}

}

func TestRectangle(t *testing.T) {
	factory := RectangleFactory{}
	shape := factory.CreateShape()
	if shape != nil {
		shape.Draw()
	}

}
