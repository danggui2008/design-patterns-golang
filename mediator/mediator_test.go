package mediator

import "testing"

func TestMediator(t *testing.T) {
	mediator := NewMediator()
	colleague1 := NewColleague("码海悔道", mediator)
	colleague2 := NewColleague("小玉", mediator)
	colleague3 := NewColleague("小李", mediator)

	mediator.Register(colleague1)

	mediator.Register(colleague2)

	mediator.Register(colleague3)

	colleague1.SendMessage("月色真好")
	colleague2.SendMessage("来了")
}
