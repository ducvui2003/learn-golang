package main

import "fmt"

func main() {

	for_normal()
	for_continued()
	for_like_while()
}

func for_normal() {
	for sum := 1; sum < 10; sum++ {
		fmt.Println(sum)
	}
}

func for_continued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func for_like_while() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
