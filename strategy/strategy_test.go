package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	calculator := NewCalculator(&Addition{})
	result := calculator.PerformOperation(20.2, 550.20)
	fmt.Printf("result:%0.2f\n", result)

	calculator.SetOperation(&Subtraction{})
	result = calculator.PerformOperation(99.90, 88.52)
	fmt.Printf("result:%0.2f\n", result)

	calculator.SetOperation(&Multiplication{})
	result = calculator.PerformOperation(100.00, 88.52)
	fmt.Printf("result:%0.2f\n", result)
}
