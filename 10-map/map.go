package main

import "fmt"

// maps -> key value pairs
func main(){
	p:= make(map[string]string)

	//add key value pairs
	p["name"] = "John"
	p["age"] = "30"
	p["city"] = "New York"

	fmt.Println(p)

	//get value by key
	fmt.Println(p["name"], p["age"], p["city"])
}