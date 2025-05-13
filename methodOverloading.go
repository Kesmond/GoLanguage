package main

import (
	"fmt"
	"time"
)

func sumTwo(x, y int) int {
	return x + y
}

func sumThree(x, y, z int) int {
	return x + y + z
}

func methodOverloadMain() {
	a := 10
	b := 10
	c := 10
	fmt.Printf("a = %v\nb = %v\nc = %v\n", a, b, c)
	fmt.Println("Go doesn't support overloading. Must use different function name!\n")
	time.Sleep(2 * time.Second)
	fmt.Println("Sum of a and b: sumTwo(a,b):", sumTwo(a, b))
	fmt.Println("Sum of a, b, c: sumThree(a,b,c):", sumThree(a, b, c))
}
