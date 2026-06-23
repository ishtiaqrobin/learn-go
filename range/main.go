package main

import "fmt"

func main() {
	// map is a dynamic array
	// myMap := map[string]string{
	// 	"name":    "Ishtiaq",
	// 	"success": "ok",
	// }

	// for _, value := range myMap {
	// 	fmt.Println(value)
	// }

	// array is a static array
	// myArr := [3]string{
	// 	"red",
	// 	"green",
	// 	"blue",
	// }

	// for i, value := range myArr {
	// 	fmt.Println(i, value)
	// }

	// slice is a dynamic array (memory allocation is done at runtime)
	// colors := []string{
	// 	"red",
	// 	"green",
	// 	"blue",
	// }

	// for i, color := range colors {
	// 	fmt.Println(i, color)
	// }

	// string is a dynamic array (underline the bytes of the string)
	name := "Next Level"

	var byteSlice = []byte(name)
	fmt.Println(byteSlice)

	// for i, value := range name {
	// 	fmt.Println(i, value)
	// }
}

// array
// slice
// string
// map
// channel
