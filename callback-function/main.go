package main

import "fmt"

// callback function
// first class citizen

// func process(sayHello func()) {
// 	sayHello()
// }

func calculator(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func main() {
	// greet := func() {
	// 	fmt.Println("Hello World!")
	// }
	// process(greet)

	add := func(n1 int, n2 int) int {
		return n1 + n2
	}

	multiply := func(n1 int, n2 int) int {
		return n1 * n2
	}

	fmt.Println(calculator(10, 20, add))
	fmt.Println(calculator(10, 20, multiply))

	// anonymous function
	result := calculator(3, 5, func(x int, y int) int {
		return x + y
	})

	fmt.Println(result)
}
