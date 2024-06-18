package memento

/*
	备忘录模式
*/
//Memento【Memento（备忘录）】
type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

// Originator【Originator（发起人】

type Originator struct {
	state string
}

func NewOriginator(state string) *Originator {
	return &Originator{state: state}
}

func (o *Originator) SetState(state string) {
	o.state = state
}

//以当前状态创建备忘录
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

//根据备忘录来恢复状态
func (o *Originator) RestoreMemento(m Memento) {
	o.state = m.state
}

//Caretaker【Caretaker（负责人）】
type Caretaker struct {
	mementos []Memento
}

func NewCaretaker() *Caretaker {
	return &Caretaker{mementos: make([]Memento, 0)}
}

func (c *Caretaker) AddMemento(memento Memento) {
	c.mementos = append(c.mementos, memento)
}

func (c *Caretaker) GetMemento(index int) (Memento, bool) {
	var memento Memento
	if index < len(c.mementos) {
		memento = c.mementos[index]
		return memento, true
	}
	return memento, false
}
