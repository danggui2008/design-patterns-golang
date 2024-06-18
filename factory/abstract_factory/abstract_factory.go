package abstractfactory

import "fmt"

/*
抽象工厂模式
*/

//操作系统【抽象产品（Abstract Product）】
type OperatingSystem interface {
	Run()
}

//Windows操作系统【具体产品（Concrete Product）】
type WindowsOS struct{}

func (windows *WindowsOS) Run() {
	fmt.Println("running in windows os ")
}

//Linux操作系统【具体产品（Concrete Product）】
type LinuxOS struct{}

func (linux *LinuxOS) Run() {
	fmt.Println("running in linux os")
}

type Application interface {
	Open()
}

//word应用程序【具体产品（Concrete Product）】
type WordApplication struct{}

func (word *WordApplication) Open() {
	fmt.Println("打开word应用程序")
}

//excel应用程序【具体产品（Concrete Product）】
type ExcelApplication struct{}

func (excel *ExcelApplication) Open() {
	fmt.Println("打开Excel应用程序")
}

//软件工厂【抽象工厂（Abstract Factory）】
type SoftwareFactory interface {
	CreateOperatingSystem() OperatingSystem
	CreateApplication() Application
}

//Window工厂【具体工厂（Concrete Factory）】
type WindowsSoftwareFactory struct{}

func (windows *WindowsSoftwareFactory) CreateOperatingSystem() OperatingSystem {
	return &WindowsOS{}
}

func (windows *WindowsSoftwareFactory) CreateApplication() Application {
	return &WordApplication{}
}

//Linux工厂【具体工厂（Concrete Factory）】
type LinuxSoftwareFactory struct{}

func (linux *LinuxSoftwareFactory) CreateOperatingSystem() OperatingSystem {
	return &LinuxOS{}
}

func (linux *LinuxSoftwareFactory) CreateApplication() Application {
	return &ExcelApplication{}
}
