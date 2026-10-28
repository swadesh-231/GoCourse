package main

import "fmt"

// by value
// func chnageNum(num int){
// 	num = 5
// 	fmt.Println("In changeNum function", num)
// }

// by reference
func chnageNumReference(num *int){
	*num = 5
	fmt.Println("In changeNum function", *num)
}

func main(){
	num := 10
	chnageNumReference(&num)
	fmt.Println("In the main function", num)
	// fmt.Println("Memory address of num", &num)
}

