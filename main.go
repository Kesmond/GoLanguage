package main

import "fmt"

func main() {
	//Declare choice as int, to take input
	var choice int

	//Main Menu
	fmt.Println("Welcome to Go Programming Language!\n")
	fmt.Println("1. Concurrency using Goroutines")
	fmt.Println("2. Built-in Error Handling")
	fmt.Println("3. No Function/Method Overload")
	fmt.Println("4. No Implicit Type Conversion")
	fmt.Print("Choose an aspect to review (1-4): ")
	fmt.Scan(&choice) //Get input of a number
	fmt.Println("")
	//Get choice variable
	switch choice {
	case 1: //If choice is 1, run goroutineMain() function
		goroutineMain()
	case 2: //If choice is 2, run errorHandlingMain() function
		errorHandlingMain()
	case 3: //If choice is 3, run methodOverloadMain() function
		methodOverloadMain()
	case 4: //If choice is 4, run typeConversionMain() function
		typeConversionMain()
	default: //Any other values, run a notice
		fmt.Println("Invalid choice")
	}
}
