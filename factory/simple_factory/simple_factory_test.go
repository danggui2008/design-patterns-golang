package simplefactory

import "testing"

func TestCircle(t *testing.T) {
	shape, _ := NewShape("Circle")
	if shape != nil {
		shape.Draw()
	}

}

func TestRectangle(t *testing.T) {
	shape, _ := NewShape("Rectangle")
	if shape != nil {
		shape.Draw()
	}

}
