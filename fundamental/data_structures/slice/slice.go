package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sliceFromArr := arr[2:3]
	fmt.Println(sliceFromArr)

	// Append element
	sliceFromArr = append(sliceFromArr, 6)
	fmt.Println(sliceFromArr)

	// Create slice from
	create_slice_from_array()
	copy_slice()
	slice_2_d()
}

// Copy slice
// Copy all element from sliceFormArr into newSlice
func copy_slice() {
	arr := [5]int{1, 2, 3, 4, 5}
	sliceFromArr := arr[2:3]
	newSlice := []int{2, 3}
	fmt.Println("newSlice before", newSlice)
	copy(newSlice, sliceFromArr)
	fmt.Println("newSlice after", newSlice)
}

func create_slice_from_array() {
	arr := [5]int{1, 2, 3, 4, 5}
	sliceFromArr := arr[0:3]
	fmt.Println("arr", arr)
	fmt.Println("Slice", sliceFromArr)
	fmt.Println("Change slice")
	arr[0] = 10
	fmt.Println("Slice", sliceFromArr)
}

func length_capacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slice_2_d() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
