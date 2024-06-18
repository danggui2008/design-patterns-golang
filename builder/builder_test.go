package builder

import (
	"fmt"
	"testing"
)

func TestHouseBuilder(t *testing.T) {
	house := House{}
	houseBuilder := ConcreteHouseBuilder{&house}

	director := NewDirector(&houseBuilder)
	director.Build()
	fmt.Printf("%#v\n", house)
}

func TestLuxuryHouseBuilder(t *testing.T) {
	house := House{}
	houseBuilder := LuxuryHouseBuilder{&house}

	director := NewDirector(&houseBuilder)
	director.Build()
	fmt.Printf("%#v\n", house)
}
