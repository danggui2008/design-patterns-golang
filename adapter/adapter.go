package adapter

/*
适配器模式
*/
import (
	"fmt"
)

// LegacyRectangle【源角色（Adaptee】
type LegacyRectangle struct{}

// 原有接口
func (r *LegacyRectangle) Display(x1 int32, y1 int32, x2 int32, y2 int32) {
	fmt.Printf("LegacyRectangle(%d,%d,%d,%d,)", x1, y1, x2, y2)
}

// Shape接口【目标角色（Target）】
type Shape interface {
	Draw(x int32, y int32, width int32, height int32)
}

//RectangleAdapter【适配器（Adapter）】

type RectangleAdapter struct {
	rectangle LegacyRectangle
}

func (r *RectangleAdapter) Draw(x int32, y int32, width int32, height int32) {
	x2 := x + width
	y2 := y + height
	r.rectangle.Display(x, y, x2, y2)
}
