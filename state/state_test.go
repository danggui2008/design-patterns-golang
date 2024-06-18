package state

import "testing"

func TestState(t *testing.T) {
	elevator := NewElevator()
	//设置电梯为开门状态
	elevator.SetState(&OpenState{})
	elevator.OpenDoor()  //开门
	elevator.Move()      //开门无法移动
	elevator.CloseDoor() //关门
	elevator.Move()      //移动
	elevator.Stop()      //停止
	elevator.OpenDoor()  //开门
}
