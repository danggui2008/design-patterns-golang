package state

import "fmt"

/*
	状态模式
*/

//ElevatorState【抽象状态角色（State）】
type ElevatorState interface {
	openDoor()
	closeDoor()
	move()
	stop()
}

//ElevatorState电梯开门状态【具体状态角色（ConcreteState）】
type OpenState struct{}

func (o *OpenState) openDoor() {
	fmt.Println("door is already opened")
}

func (o *OpenState) closeDoor() {
	fmt.Println("closing the door")
}
func (o *OpenState) move() {
	fmt.Println("cannot move while is the door is opened")
}
func (o *OpenState) stop() {
	fmt.Println("is stopped")
}

//CloseState电梯关门状态【具体状态角色（ConcreteState）】
type CloseState struct{}

func (c *CloseState) openDoor() {
	fmt.Println("opening the door")
}

func (c *CloseState) closeDoor() {
	fmt.Println("the door is closed")
}
func (c *CloseState) move() {
	fmt.Println("moving")
}
func (c *CloseState) stop() {
	fmt.Println("stopping")
}

//Elevator电梯【环境角色（Context）】
type Elevator struct {
	state ElevatorState
}

func (e *Elevator) OpenDoor() {
	e.state.openDoor()
}

func (e *Elevator) CloseDoor() {
	e.state.closeDoor()
}
func (e *Elevator) Move() {
	e.state.move()
}
func (e *Elevator) Stop() {
	e.state.stop()
}
func (e *Elevator) SetState(state ElevatorState) {
	e.state = state
}

func NewElevator() *Elevator {
	return &Elevator{state: &CloseState{}}
}
