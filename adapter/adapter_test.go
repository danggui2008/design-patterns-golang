package adapter

import "testing"

func TestAdapter(t *testing.T) {
	rectangle := LegacyRectangle{}
	adapter := RectangleAdapter{rectangle: rectangle}
	adapter.Draw(10, 20, 20, 30)
}
