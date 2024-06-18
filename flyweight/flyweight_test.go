package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	factory := NewShapeFactory()
	colors := [4]string{"red", "green", "blue", "yellow"}
	for i := 0; i < 10; i++ {
		for c := 0; c < 4; c++ {
			x := float32(i + c + 10)
			y := float32(i + c + 20)
			shape := factory.GetShape(colors[c])
			coordinate := NewCoordinate(x, y)
			shape.Draw(coordinate)
		}
	}
}
