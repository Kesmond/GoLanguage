package main

import "fmt"

func typeConversionMain() {
	var a int = 10
	var b float32 = 10.5

	fmt.Printf("a = %v, is a %T\n", a, a)
	fmt.Printf("b = %v, is a %T\n\n", b, b)

	fmt.Println("Go doesn't support implicit type conversion!")
	fmt.Println("When adding the two, a must be float")
	fmt.Println("")
	sum := float32(a) + b
	fmt.Println("a + b = ", sum)
}
