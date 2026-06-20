package main

import "fmt"

func main() {
	var numbers [6]int
	numbers[0] = 1
	numbers[1] = 5
	numbers[5] = 40
	fmt.Println(numbers)

	// fmt.Println("Length of array is :",len(numbers))

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
