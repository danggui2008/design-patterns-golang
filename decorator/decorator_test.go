package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	//普通咖啡
	simpleCoffee := SimpleCoffee{}
	fmt.Printf("花了%0.2f(元)，买了一杯%s\n", simpleCoffee.Cost(), simpleCoffee.Desc())

	//加牛奶的咖啡
	milkCoffee := NewMilkDecorator(&simpleCoffee)
	fmt.Printf("花了%0.2f(元)，买了一杯%s\n", milkCoffee.Cost(), milkCoffee.Desc())

	//给牛奶的咖啡加点糖
	sugarCoffee := NewSugarDecorator(milkCoffee)
	fmt.Printf("花了%0.2f(元)，买了一杯%s\n", sugarCoffee.Cost(), sugarCoffee.Desc())
}
