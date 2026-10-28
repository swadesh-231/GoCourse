package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Basic Calculator")
	fmt.Println("Enter q to quit")

	for {
		fmt.Print("First number: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read input")
			break
		}
		text := strings.TrimSpace(line)
		if text == "q" || text == "Q" {
			fmt.Println("Bye")
			break
		}

		a, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		fmt.Print("Operator (+ - * /): ")
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read input")
			break
		}
		op := strings.TrimSpace(line)

		fmt.Print("Second number: ")
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read input")
			break
		}
		b, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		switch op {
		case "+":
			fmt.Println("Result:", a, "+", b, "=", a+b)
		case "-":
			fmt.Println("Result:", a, "-", b, "=", a-b)
		case "*":
			fmt.Println("Result:", a, "*", b, "=", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Cannot divide by zero")
				continue
			}
			fmt.Println("Result:", a, "/", b, "=", a/b)
		default:
			fmt.Println("Unknown operator. Use + - * /")
		}
	}
}
