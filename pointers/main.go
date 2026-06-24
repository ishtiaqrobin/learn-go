package main

import "fmt"

func change(x *int) {
	*x = 500
	println("Inside function :", *x)
}

func main() {
	// a := 42

	// p depends on a variable
	// p := &a // 0x289e38d20b0 --> pointer - (memory address)

	// a = 30

	// *p = 500 // change the value of a because p points to a

	// fmt.Println("a :", a)
	// fmt.Println("p :", p)  // 0x289e38d20b0 --> pointer
	// fmt.Println("p :", *p) // pointer dereferencing - 30

	// fmt.Println("a :", &a)

	// & --> address of a variable (memory address)
	// * --> dereferencing operator (value from memory address)

	// y := 20

	// change(&y)

	// fmt.Println("Outside function :", y)

	bigArray := [5]int{10, 20, 30, 40, 50}

	// modifyWithoutPointer(bigArray)
	// fmt.Println("After without pointer :", bigArray)

	modifyWithPointer(&bigArray)
	fmt.Println("After with pointer :", bigArray)
}

func modifyWithoutPointer(arr [5]int) {
	arr[0] = 999 // it will not modify the original array because it is a copy
	fmt.Println("Inside without pointer :", arr)
}

func modifyWithPointer(arr *[5]int) {
	// * not required because go internally dereferences
	arr[0] = 777 // it will modify the original array because it is a pointer
	fmt.Println("Inside with pointer :", *arr)
}

// Pointer basic
// Dereference example
// Modify value using pointer
// Pointer in function
// Nil pointer
