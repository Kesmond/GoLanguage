package main

import (
	"fmt"
	"time"
)

// First function, adding two values
func sumTwo(x, y int) int {
	return x + y //Add the two values and return them
}

// Second function, adding three values
func sumThree(x, y, z int) int {
	return x + y + z //Add the three values and return them
}

// 'Main' function when getting a call from main.go
func methodOverloadMain() {
	//Set a,b, and c to 10
	a := 10
	b := 10
	c := 10
	//Print the values of a,b,c for user clarity
	fmt.Printf("a = %v\nb = %v\nc = %v\n", a, b, c)
	//Explanation on the bad aspect
	fmt.Println("Go doesn't support overloading. Must use different function name!\n")
	time.Sleep(2 * time.Second) //Give some time for users to read line 26

	//Print the sum of two values and which function is used
	fmt.Println("Sum of a and b: sumTwo(a,b):", sumTwo(a, b))
	//Print the sum of three values and which function is used
	fmt.Println("Sum of a, b, c: sumThree(a,b,c):", sumThree(a, b, c))
}
