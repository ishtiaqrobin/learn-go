package main

import "fmt"

func main() {
	// anonymous function
	// makeCoffee := func() {
	// 	fmt.Printf("Making Coffee")
	// }

	// makeCoffee()

	// Immediately Invoked Function Expression - IIFE
	// func() {
	// 	println("Hello World")
	// }()

	// IIFE with parameters
	func(coffeeType string) {
		fmt.Printf("Making %s coffee", coffeeType)
	}("Espresso")
}
