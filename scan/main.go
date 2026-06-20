package main

import "fmt"

func main() {
	var choice int

	fmt.Print("Enter your number: ")
	fmt.Scan(&choice)

	fmt.Println("You entered", choice)
}
