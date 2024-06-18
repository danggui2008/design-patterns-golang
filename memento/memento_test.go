package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	//发起人
	originator := NewOriginator("state 1")
	//创建备注录：state 1
	memento := originator.CreateMemento()
	//负责人
	caretaker := NewCaretaker()
	caretaker.AddMemento(*memento)
	originator.SetState("state 2")
	//state 2
	fmt.Printf("%v\n", *originator)

	//恢复状态(state 2-》state 1)
	m, _ := caretaker.GetMemento(0)
	originator.RestoreMemento(m)
	//已恢复到原来状态：state 1
	fmt.Printf("%v\n", *originator)
}
