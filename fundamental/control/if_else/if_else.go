package main

import "fmt"

func main() {
	number := 10
	if number/2 == 0 {
		fmt.Printf("%d is even number", number)
	} else {
		fmt.Printf("%d is odd number", number)
	}
}
