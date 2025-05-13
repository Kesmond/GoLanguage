package main

import "fmt"

func main() {
	var choice int

	fmt.Println("Welcome to Go Programming Language!\n")
	fmt.Println("1. Concurrency using Goroutines")
	fmt.Println("2. Error Handling in Function")
	fmt.Println("3. Function/Method Overload")
	fmt.Println("4. Bad Aspect ")
	fmt.Print("Choose an aspect to review: ")
	fmt.Scan(&choice)
	fmt.Println("")
	switch choice {
	case 1:
		goroutineMain()
	case 2:
		errorHandlingMain()
	case 3:
		methodOverloadMain()
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("Invalid choice")
	}
	//methodOverloadMain()
}
