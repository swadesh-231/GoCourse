package main

import "fmt"

func main(){
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("Almost weekend")
	default:
		fmt.Println("Weekend")
	}

	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 75:
		fmt.Println("Good")
	default:
		fmt.Println("Keep practicing")
	}

	days := []string{"Monday", "Sunday", "Wednesday", "Saturday", "Friday"}
loop:
	for _, name := range days {
		switch name {
		case "Sunday":
			fmt.Println("skip Sunday")
			continue
		case "Saturday":
			fmt.Println("stop at Saturday")
			break loop
		default:
			fmt.Println("working day:", name)
		}
		fmt.Println("after switch:", name)
	}
}
