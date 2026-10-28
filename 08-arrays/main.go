package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers)


	for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }

    // Or use range
    for index, value := range numbers {
        fmt.Println(index, value)
    }

	marks := [5]int{90, 80, 70, 60, 50}
	fmt.Println(marks)

    var cities  = [7]string{"New York", "Los Angeles", "Chicago", "Houston", "Miami", "San Francisco", "Seattle"}
    fmt.Println(cities)


    //2D arrays
    var matrix = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println(matrix)
   

}
