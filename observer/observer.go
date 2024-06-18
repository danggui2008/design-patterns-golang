package observer

import "fmt"

/*
	观察者模式
*/

//Subject【抽象主题 (Subject)】
type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

//ConcreteSubject【具体主题 (Concrete Subject)】
type ConcreteSubject struct {
	state     int32
	observers []Observer
}

func (s *ConcreteSubject) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers() {
	for _, o := range s.observers {
		o.Update(s.state)
	}
}

func (s *ConcreteSubject) SetState(state int32) {
	s.state = state
}

func NewSubject() *ConcreteSubject {
	return &ConcreteSubject{state: 0, observers: make([]Observer, 0)}
}

type Observer interface {
	Update(state int32)
}

//ConcreteObserver【具体观察者 (Concrete Observer)】
type ConcreteObserver struct {
	id int32
}

func NewObserver(id int32) Observer {
	return &ConcreteObserver{id: id}
}

func (c *ConcreteObserver) Update(state int32) {
	fmt.Printf("Observer[%d] state:%d is updated\n", c.id, state)
}
