package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func errorHandlingMain() {
	var a float64
	var b float64
	fmt.Println("Enter two numbers to be divided!")
	fmt.Print("Numerator/First number: ")
	fmt.Scan(&a)
	fmt.Print("Denominator/Second number: ")
	fmt.Scan(&b)
	fmt.Println()
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
