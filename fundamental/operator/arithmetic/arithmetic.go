package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2
	sum := num1 + num2
	fmt.Printf("%d - %d = %d\n", num1, num2, sum)

	//quotient = nguyen
	quotient := num1 / num2
	fmt.Printf("%d / %d = %d\n", num1, num2, quotient)

	// Get actual number
	num3 := 11.0
	num4 := 4.0
	actual := num3 / num4
	fmt.Printf("%g / %g = %g\n", num1, num2, actual)

	//Modulus Operator
	remainder := num1 % num2
	fmt.Printf("%d / %d = %d\n", num1, num2, remainder)

}
