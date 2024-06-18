package iterator

/*
	迭代器模式
*/
//迭代器接口
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

//具体迭代器
type ConcreteIterator[T any] struct {
	items    []T //元素
	position int //位置
}

//实现迭代器方法HasNext
func (ci *ConcreteIterator[T]) HasNext() bool {
	return len(ci.items) > ci.position
}

//实现迭代器方法Next
func (ci *ConcreteIterator[T]) Next() T {
	var value T
	if ci.HasNext() {
		value = ci.items[ci.position]
		ci.position += 1
	}
	return value

}

//创建迭代器接口
type IterableCollection[T any] interface {
	CreateIterator() Iterator[T]
}

//集合
type ConcreteCollection[T any] struct {
	items []T
}

func NewConcreteCollection[T any]() *ConcreteCollection[T] {
	return &ConcreteCollection[T]{items: make([]T, 0)}
}

//创建迭代器
func (c *ConcreteCollection[T]) CreateIterator() Iterator[T] {
	return &ConcreteIterator[T]{
		items:    c.items,
		position: 0,
	}
}

//添加元素
func (c *ConcreteCollection[T]) Add(item T) {
	c.items = append(c.items, item)
}
