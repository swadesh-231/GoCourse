package main

import "fmt"

func main(){
	// numbers := []int{10, 20, 30, 40, 50}

    // for index, value := range numbers {
    //     fmt.Println(index, value)
    // }

	d := map[string]string{
		"name": "John",
		"age": "30",
		"city": "New York",
	}

	for key, value := range d {
		fmt.Println(key, value)
	}

	for _, value := range d {
		fmt.Println(value)
	}
}