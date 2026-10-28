package main

import "fmt"

func main(){
	greet()

	sayHello("Swadesh")

	result := add(10, 20)
	fmt.Println(result)

	difference := subtract(20, 8)
	fmt.Println(difference)

	sum, product := calculate(4, 5)
	fmt.Println(sum, product)

	area := rectangle(5, 3)
	fmt.Println(area)

	total := addMany(1, 2, 3, 4)
	fmt.Println(total)

	multiply := func(a int, b int) int {
		return a * b
	}
	fmt.Println(multiply(6, 7))
}

func greet() {
	fmt.Println("Hello")
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func calculate(a int, b int) (int, int) {
	return a + b, a * b
}

func rectangle(length int, width int) (area int) {
	area = length * width
	return
}

func addMany(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total = total + n
	}
	return total
}
