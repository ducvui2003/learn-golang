package main

import (
	"fmt"
	"strings"
)

func main() {
	var message string
	message = "hello world"

	fmt.Println(message)

	sliceStr := []string{"hello", "world"}
	fmt.Println(strings.Join(sliceStr, " "))
}
