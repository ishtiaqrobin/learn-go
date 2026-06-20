package main

import "fmt"

// Variable scope
func main() {
	sugar := 100

	makeCoffee := func() {
		// Variable shadowing
		// sugar := 50

		// Modifying outer variable (when use = keyword not use := keyword)
		sugar = 50

		coffee := "Latte"
		fmt.Printf("Making a cup of %s coffee with %d grams of sugar \n", coffee, sugar)
		fmt.Printf("Value of inner sugar variable: %d\n", sugar)
	}

	makeCoffee()

	fmt.Printf("Value of outer sugar variable: %d\n", sugar)
}
