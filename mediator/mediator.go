package mediator

import "fmt"

/*
	中介者模式
*/
//ChatMediator【Mediator（抽象中介者）】
type ChatMediator interface {
	//通知接口：同事通过此接口与同事通信
	Notify(name string, message string)
	Register(c Colleague)
}

// Colleague【Colleague（抽象同事类）】
type Colleague interface {
	getName() string
	//同事行动起来：中介者与之通信
	Action(message string)
	//发送消息：通过中介者与与具体同事通信
	SendMessage(message string)
}

//ConcreteChatMediator【ConcreteMediator（具体中介者）】
type ConcreteChatMediator struct {
	colleagues map[string]Colleague
}

func (m *ConcreteChatMediator) Register(c Colleague) {
	m.colleagues[c.getName()] = c
}

func (m *ConcreteChatMediator) Notify(name string, message string) {
	c, ok := m.colleagues[name]
	if ok {
		c.Action(message)
	}
}

func NewMediator() ChatMediator {
	return &ConcreteChatMediator{colleagues: make(map[string]Colleague)}
}

//ConcreteColleague【ConcreteColleague（具体同事类）】
type ConcreteColleague struct {
	name     string
	mediator ChatMediator
}

func NewColleague(name string, mediator ChatMediator) Colleague {
	return &ConcreteColleague{name: name, mediator: mediator}
}

func (c *ConcreteColleague) getName() string {
	return c.name
}

func (c *ConcreteColleague) Action(message string) {
	fmt.Printf("[%s]收到消息:%s，我马上行动起来\n", c.name, message)
}

func (c *ConcreteColleague) SendMessage(message string) {
	fmt.Printf("将要给[%s]发送消息:%s\n", c.name, message)
	c.mediator.Notify(c.name, message)
}
