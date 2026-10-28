package main

import "fmt"

func main(){
	total := sum(10, 20, 30, 40)
	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total = total + n
	}
	return total
}
