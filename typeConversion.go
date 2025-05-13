package main

import "fmt"

// 'Main' function when getting a call from main.go
func typeConversionMain() {
	//Declare two variables, int and float32 variables
	var a int = 10
	var b float32 = 10.5

	//Print the values and data types, for clarity
	fmt.Printf("a = %v, is a %T\n", a, a)
	fmt.Printf("b = %v, is a %T\n\n", b, b)

	//Print the explanation of the bad aspect
	fmt.Println("Go doesn't support implicit type conversion!")
	fmt.Println("When adding the two, a must be float")
	fmt.Println("")
	//Operations can't be done on different variables
	//Changing a to float and add with b, store it in 'sum' variable
	sum := float32(a) + b
	fmt.Println("a + b = ", sum) //Print the sum value
}
