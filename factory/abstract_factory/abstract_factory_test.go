package abstractfactory

import "testing"

func TestWindowsOS(t *testing.T) {
	factory := WindowsSoftwareFactory{}
	operating := factory.CreateOperatingSystem()
	application := factory.CreateApplication()
	operating.Run()
	application.Open()
}

func TestLinuxOS(t *testing.T) {
	factory := LinuxSoftwareFactory{}
	operating := factory.CreateOperatingSystem()
	application := factory.CreateApplication()
	operating.Run()
	application.Open()
}
