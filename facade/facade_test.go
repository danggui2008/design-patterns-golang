package facades

import "testing"

func TestFacade(t *testing.T) {
	homeTheater := NewHomeTheater(StereoSystem{}, Projector{}, LightsControl{})
	homeTheater.WatchMovie()
	homeTheater.EndMove()

}
