package main

import "fmt"

//* defer version: 1.0

// func deferred(result int) {
// 	fmt.Println("deferred result :", result)
// }

// func example() int {
// 	result := 10

// 	defer deferred(result)

// 	result = 20

// 	fmt.Println("Inside example:", result)
// 	return result
// }

// func main() {
// 	example()
// }

//? Result:
// Inside example: 20
// deferred result : 10

//* deferred version: 2.0

// func example() int {
// 	result := 10

// 	defer func() {
// 		fmt.Println("deferred result :", result) // 60
// 	}()

// 	fmt.Println("Inside example:", result) // 10

// 	result = result + 50 // 60
// 	return result        // 60
// }

// func main() {
// 	fmt.Println("Return result:", example()) // 60
// }

//? Result:
// Inside example: 10
// deferred result : 60
// Return result: 60

//* deferred version: 2.1

// func example() int {
// 	result := 10

// 	defer func() {
// 		result = result + 200                    // 310
// 		fmt.Println("deferred result :", result) // 310
// 	}()

// 	fmt.Println("Inside example:", result) // 10

// 	result = result + 100 // 110
// 	return result         // return 110
// }

// func main() {
// 	fmt.Println("Return result:", example()) // 110
// }

//? Result:
// Inside example: 10
// deferred result : 310
// Return result: 110

//* deferred version: 3.0
//* named return values

func example() (result int) { // zero value (0)
	result = 10 // reassign

	defer func() {
		result = result + 200                    // 310
		fmt.Println("deferred result :", result) // 310
	}()

	// follow LIFO
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(2)
	defer fmt.Println(4)

	fmt.Println("Inside example:", result) // 10

	result = result + 100 // 110
	return                // 0x485468191
}

func main() {
	fmt.Println("Return result:", example()) // 310
}

// Inside example: 10
// deferred result : 310
// Return result: 310

// * example of real life cycle
func connectDB() {
	defer fmt.Println("Closing database connection...")
	defer fmt.Println("Connecting to database...")
	// some operations
}
