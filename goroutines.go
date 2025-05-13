package main

import (
	"fmt"
	"time" //Use to delay
)

// Function without 'go' command
func regularFunction() {
	//Gives clarity that the function has run
	fmt.Println("Entered regularFunction()")
	//Delays process for 10 sec, so goroutineFunction can finish its process
	time.Sleep(10 * time.Second)
}

// Function with 'go' command
func goroutineFunction() {
	//Gives clarity that the function has run
	fmt.Println("Entered goroutineFunction()")
	time.Sleep(5 * time.Second)               //Delays process for 5 sec
	fmt.Println("goroutineFunction finished") //Clarity that the function has finished
}

// 'Main' function when getting a call from main.go
func goroutineMain() {
	go goroutineFunction() //Starting the first function using 'go' command

	//Clarity that the line below has been execute
	fmt.Println("Main, Line below go goroutineFunction()")

	regularFunction() //Calling the second function without 'go' command

	//Clarity that the line below has been execute
	fmt.Println("Main, Line below regularFunction()")
}
