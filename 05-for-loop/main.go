package main

import "fmt"

func main(){
	for i := 1; i <= 5; i++ {
		fmt.Println("count:", i)
	}

	n := 1
	for n <= 3 {
		fmt.Println("while-style:", n)
		n++
	}

	fruits := []string{"apple", "banana", "mango"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}
}
