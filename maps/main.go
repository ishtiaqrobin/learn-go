package main

import "fmt"

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	// myMap := make(map[string]int)
	// myMap := make(map[int]string)

	// myMap["user1Score"] = 50
	// myMap[3] = "two"

	// fmt.Println(myMap["user1Score"])
	// fmt.Println(myMap[3])

	// myMap := map[string]string{
	// 	"name":    "Ishtiaq",
	// 	"success": "ok",
	// }

	// delete(myMap, "name")

	// fmt.Println(myMap)

	// fmt.Println(myMap["name"])

	myMap := map[string]user{
		"data": user{
			name:       "Ishtiaq",
			age:        23,
			isLoggedIn: true,
		},
	}

	fmt.Println(myMap)

}

// Creating map (using make and literal)
// Accessing map values
// Adding new key-value pair to map
// Deleting key-value pair from map
// Map of struct
// Iterating over map
