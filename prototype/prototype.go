package prototype

/*
	原型模式
*/
type Cloneable interface {
	Clone() *Shape
}

type Shape struct {
	sType string
}

func (s *Shape) Clone() *Shape {
	return &Shape{s.sType}
}
