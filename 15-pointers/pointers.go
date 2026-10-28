package main

import "fmt"

// by value
func chnageNum(num int){
	num = 5
	fmt.Println("in changeNum", num)
}


func main(){
	num := 10
	chnageNum(num)
	fmt.Println("In the main function", num)
}

