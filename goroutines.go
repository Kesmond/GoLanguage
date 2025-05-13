package main

import (
	"fmt"
	"time"
)

func regularFunction() {
	fmt.Println("Entered regularFunction()")
	time.Sleep(10 * time.Second)
}

func goroutineFunction() {
	fmt.Println("Entered goroutineFunction()")
	time.Sleep(5 * time.Second)
	fmt.Println("goroutineFunction finished")
}

func goroutineMain() {
	go goroutineFunction()
	fmt.Println("Main, Line below go goroutineFunction()")
	regularFunction()
	fmt.Println("Main, Line below regularFunction()")
}
