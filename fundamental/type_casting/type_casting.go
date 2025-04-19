package main

import "fmt"

func main() {
	// explicit
	var float1 float64 = 12.3
	var int1 int = int(float1)
	fmt.Println(int1)

	fmt.Printf("Float Value is %g\n", float1)
	fmt.Printf("Integer Value is %d", int1)

	// Implicit
	// Wrong
	// var number int = 4.34

	// fmt.Printf("Number is %g", number)
}
