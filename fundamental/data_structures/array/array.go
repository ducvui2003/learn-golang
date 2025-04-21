package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	languages := [3]string{"Go", "Java", "Javascript"}
	fmt.Println(languages[1])

	languages[2] = "Kotlin"
	lengthOfLanguages := len(languages)
	fmt.Print(lengthOfLanguages)

	// Loop each element in array
	for i := 0; i < len(languages); i++ {
		fmt.Println(languages[i])
	}

	for index, element := range languages {
		fmt.Printf("Index %d: %s \n", index, element)
	}

	// Array 2 dimension
	array2d := [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < len(array2d); i++ {
		for j := 0; j < len(array2d[i]); j++ {
			fmt.Println(array2d[i][j])
		}
	}

}
