package bridge

import "testing"

func TestBridge(t *testing.T) {
	red := Red{}
	circle := Circle{color: &red}
	circle.Draw()

	blue := Blue{}
	square := Square{color: &blue}
	square.Draw()
}
