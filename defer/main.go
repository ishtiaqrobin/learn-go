package main

import "fmt"

func main() {
	defer fmt.Println("I am deferred") // deferred function

	fmt.Println("I am from the main function")
}
