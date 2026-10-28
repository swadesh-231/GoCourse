package main

import "fmt"

func main(){
	score := 75

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	age := 20
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote yet")
	}
}
