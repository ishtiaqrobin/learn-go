package main

import "fmt"

//* Variadic function

func add(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		// total += number
		total = total + number
	}

	return total
}

func greet(prefix string, mps ...string) {
	for _, mp := range mps {
		fmt.Println(prefix, mp)
	}
}

func main() {
	// sum := add(5, 10, 25, 10)
	// fmt.Println(sum)

	// greet("Welcome", "Jamal", "Kamal", "Khairul")

	mps := []string{"Jamal", "Kamal", "Khairul"}
	greet("Welcome", mps...) // Jamal, Kamal, Khairul
}

// flexible amount of argument
// must be the last parameter
// internally slice
