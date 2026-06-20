package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Printf("Make coffee %d \n", coffeeNo)
}

func main() {
	// for initialization; condition; increment/decrement (post statement)
	// for i := 1; i <= 5; i++ {
	// 	makeCoffee(i)
	// }

	// while style for loop
	// i := 1
	// for i <= 5 {
	// 	makeCoffee(i)
	// 	i++
	// }

	// infinite loop
	// for {
	// 	makeCoffee(1)
	// }

	// break (break the loop), continue (skip the iteration)
	// for i := 1; i <= 10; i++ {
	// 	if i == 6 {
	// 		break
	// 	}

	// 	makeCoffee(i)
	// }

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip the iteration and move to the next iteration
		}

		makeCoffee(i)
	}

}

// i := 1, true, run the body, increment
// i := 2, true, run the body, increment
// i := 3, true, run the body, increment
// i := 4, true, run the body, increment
// i := 5, true, run the body, increment
// i := 6, false, break, no increment
