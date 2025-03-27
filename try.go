package main

import (
	"fmt"
	"time"
)

func main() {

	//Handles concurrency
	go display("Hello, Goroutine!")
	display("Hello, Main!")

	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Println(s, i)
			time.Sleep(500 * time.Millisecond)
		}
	}("Hello from Anonymous Gourutine!")

	time.Sleep(2 * time.Second)
	fmt.Println("Main function complete.")

	//Executes methodOverloading.go
	methodMain()
}

func ProcessNumber(input int, output chan<- int) {
	go func() {
		time.Sleep(2 * time.Second)
		output <- input * 2
	}()
}

func display(str string) {
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str, i)
	}
}
