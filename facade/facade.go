package facades

import "fmt"

/*
	门面模式
*/

//StereoSystem音响【子系统角色（SubSystem）】
type StereoSystem struct{}

func (s *StereoSystem) turnOn() {
	fmt.Println("打开音响")
}

func (s *StereoSystem) turnOff() {
	fmt.Println("关闭音响")
}

// Projector投影仪【子系统角色（SubSystem）】
type Projector struct{}

func (p *Projector) turnOn() {
	fmt.Println("打开投影仪")
}

func (p *Projector) turnOff() {
	fmt.Println("关闭投影仪")
}

// LightsControl灯光控制【子系统角色（SubSystem）】
type LightsControl struct{}

func (l *LightsControl) turnOn() {
	fmt.Println("灯光已打开")
}

func (l *LightsControl) turnOff() {
	fmt.Println("关闭灯光")
}

// HomeTheaterFacade家庭影院外观【外观角色（Facade）】
type HomeTheaterFacade struct {
	stereo    StereoSystem
	projector Projector
	lights    LightsControl
}

func NewHomeTheater(stereo StereoSystem, projector Projector, lights LightsControl) *HomeTheaterFacade {
	return &HomeTheaterFacade{stereo: stereo, projector: projector, lights: lights}
}

func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("开始看电影")
	h.lights.turnOff()
	h.projector.turnOn()
	h.stereo.turnOn()
}

func (h *HomeTheaterFacade) EndMove() {
	fmt.Println("电影结束")
	h.stereo.turnOff()
	h.projector.turnOff()
	h.lights.turnOn()
}
