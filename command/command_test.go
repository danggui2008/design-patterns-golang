package command

import "testing"

func TestCommand(t *testing.T) {
	//客户端（Client）
	//命令接收者
	light := Light{}
	//开灯命令
	lightOnCommand := NewLightOnCommand(&light)
	//关灯命令
	lightOffCommand := NewLightOffCommand(&light)
	//命令调用者：远程控制1
	remoteControl := RemoteControl{}
	//设置开灯
	remoteControl.SetCommand(lightOnCommand)

	remoteControl.PressButton()
	//设置关烟命令
	remoteControl.SetCommand(lightOffCommand)

	remoteControl.PressButton()
}
