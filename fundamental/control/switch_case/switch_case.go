package main

import (
	"fmt"
	"time"
)

func main() {
	dayOfWeek := 3
	// switch statement auto terminate if one case match (not need break)
	switch dayOfWeek {

	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
		fallthrough // if case 3 match, it will run case 4 and terminate

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")
	}

	switch_case_no_condition()
	switch_case_multi_cases()
}

func switch_case_no_condition() {
	t := time.Now()
	// switch no condition = switch true
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func switch_case_multi_cases() {
	dayOfWeek := "Sunday"

	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}
}
