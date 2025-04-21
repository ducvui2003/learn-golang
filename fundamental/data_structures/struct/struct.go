package main

import "fmt"

type Rectangle func(int, int) int

type Vertex struct {
	X int
	Y int

	rect Rectangle
}

func main() {
	// Can use func like a field of struct
	v := Vertex{X: 1, Y: 2, rect: func(length int, breadth int) int {
		return length * breadth
	}}

	v.X = 4
	fmt.Println(v.X)
}
