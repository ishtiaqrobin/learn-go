package main

import "fmt"

func main() {
	// var numbers = [6]int{10, 20, 50, 60} // partial array initialization
	var orders = [6]int{10, 20, 30, 40, 50, 60}

	// [start:end(excluded)] len, cap, pointer
	slice1 := orders[2:4]
	fmt.Println(slice1)

	// slice1[0] = 100

	slice1 = append(slice1, 500)
	slice1 = append(slice1, 600)
	slice1 = append(slice1, 700)
	slice1 = append(slice1, 800)

	// fmt.Println(slice1)
	// fmt.Println(orders)

	fmt.Println("The length of the slice is", len(slice1))
	fmt.Println("The capacity of the slice is", cap(slice1))
}
