package command

import "fmt"

/*
	命令模式
*/

//Command命令【命令（Command）】
type Command interface {
	execute()
}

//LightOnCommand开灯命令【具体命令（Concrete Command）】
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) Command {
	return &LightOnCommand{light: light}
}

//实现命令接口
func (lo *LightOnCommand) execute() {
	lo.light.turnOn()
}

//LightOffCommand关灯命令【具体命令（Concrete Command）】
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) Command {
	return &LightOffCommand{light: light}
}

//实现命令接口
func (lo *LightOffCommand) execute() {
	lo.light.turnOff()
}

//Light灯【接收者（Receiver）】
type Light struct{}

func (l *Light) turnOn() {
	fmt.Println("Light is on")
}

func (l *Light) turnOff() {
	fmt.Println("Light is off")
}

//RemoteControl【调用者/请求者（Invoker）】
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	if r.command != nil {
		r.command.execute()
	}
}
