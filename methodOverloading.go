package main

import "fmt"

func sumTwoInt(x, y int) int {
	return x + y
}

func sumThreeInt(x, y, z int) int {
	return x + y + z
}

func methodMain() {
	x := 10
	y := 10
	z := 10
	fmt.Println("Two Integers: ", sumTwoInt(x,y))
	fmt.Println("Three Integers: ", sumThreeInt(x,y,z))
}