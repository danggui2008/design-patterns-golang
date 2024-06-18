package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	observer1 := NewObserver(1)
	observer2 := NewObserver(2)
	observer3 := NewObserver(3)
	subject := NewSubject()
	subject.AddObserver(observer1)
	subject.AddObserver(observer2)
	subject.AddObserver(observer3)

	subject.SetState(100)
	subject.NotifyObservers()
	fmt.Println("remove observer2")
	subject.RemoveObserver(observer2)

	subject.SetState(200)
	subject.NotifyObservers()
}
