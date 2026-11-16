package main

import (
	"fmt"
	"myproject/calculator"
)

func main() {
	result := calculator.Add(10, 20)

	fmt.Println("Addition:", result)

	result2 := calculator.Subtract(20, 10)

	fmt.Println("Subtraction:", result2)
}
