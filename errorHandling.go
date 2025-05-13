package main

import (
	"errors"
	"fmt"
)

// Function to divide two numbers
func divide(a, b float64) (float64, error) {
	if b == 0 { //If b is zero
		//return the value 0 of the division and an error message
		return 0, errors.New("division by zero is not allowed")
	}
	//If b isn't a 0, return a divided by b and no error message
	return a / b, nil
}

// 'Main' function when getting a call from main.go
func errorHandlingMain() {
	//Declare two float variables for calculation
	var a float64
	var b float64

	//Receiving an input of two integers
	fmt.Println("Enter two numbers to be divided!")
	fmt.Print("Numerator/First number: ")
	fmt.Scan(&a)
	fmt.Print("Denominator/Second number: ")
	fmt.Scan(&b)

	//Run divide function with a and b as parameters, and return two values
	//first variable as the result of the division stored in 'result' variable
	//second variable as the error message stored in the 'err' variable
	result, err := divide(a, b)
	if err != nil { //If err has values in it
		fmt.Println("Error:", err) //Print the error message
	} else {
		fmt.Println("Result:", result) //Print the result of the division
	}
}
