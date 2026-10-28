package main

import "fmt"

func main(){
	//slices are dynamic arrays

	// numbers := []int{10, 20, 30, 40, 50}
	// fmt.Println(numbers)

	// cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Miami"}
	// fmt.Println(cities)

	// var marks[] int;
	// marks = append(marks, 90, 80, 70, 60, 50)
	// fmt.Println(marks)


	//length and capacity of a slice
	// make([]int, length, capacity)

	numbers := make([]int, 1, 3) // 3 is the length and 5 is the capacity
	numbers = append(numbers, 10, 20, 30, 40, 50)
	fmt.Println(cap(numbers)) 
    fmt.Println(numbers)     
    fmt.Println(len(numbers)) 

	//slice operators 
	ages := []int{10, 20, 30, 40, 50}

    fmt.Println(ages[1:4])


}