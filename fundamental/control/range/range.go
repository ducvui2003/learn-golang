package main

import "fmt"

func main() {
	range_array()
	range_with_string()
	range_with_map()
	range_access_key()
}

func range_array() {
	numbers := [5]int{1, 2, 3, 4, 5}
	for index, element := range numbers {
		fmt.Print(index, element)
	}
}
func range_with_string() {
	string := "Golang"
	fmt.Println("Index Character: ")

	for i, item := range string {
		fmt.Printf("%d= %c \n", i, item)
	}
}
func range_with_map() {
	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	fmt.Println("Marks obtained:")

	for key, value := range subjectMarks {
		fmt.Println(key, ":", value)
	}
}
func range_access_key() {
	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	fmt.Println("Marks obtained:")

	for value := range subjectMarks {
		fmt.Println(value)
	}
}
