package main

import "fmt"

func main() {
	basic()
}

func basic() {
	students := map[int]string{1: "Join", 2: "Ken"}
	fmt.Println(students)

	// add element
	students[0] = "Ken"
	fmt.Println(students)

	// Delete element
	delete(students, 1)
	fmt.Println(students)

	// Loop
	for key, value := range students {
		fmt.Println(key, value)
	}

	// Create by make
	// create a map using make()
	student := make(map[int]string)

	// add elements to the map
	student[1] = "Harry"
	student[2] = "Lilly"
	student[5] = "Harmonie"
}
