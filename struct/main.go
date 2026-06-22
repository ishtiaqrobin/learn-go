package main

import "fmt"

type additionalInfo struct {
	phone   int
	address string
}

// type user struct {
// 	name  string
// 	email string

// 	metaInfo additionalInfo
// }

type user struct {
	name string
	age  int
	role string
}

func main() {

	// john := user{"John Doe", "johndoe@go.dev"} // Positional parameters

	// Key-Value parameters
	// john := user{
	// 	name:  "John Doe",
	// 	email: "johndoe@go.dev",
	// }

	// john.name = "Jane"

	// fmt.Println(john)
	// fmt.Printf("%+v", john) // %+v will print the struct
	// fmt.Println(john.name)

	// var user1 user

	// user1.name = "John Doe"
	// user1.email = "johndoe@go.dev"

	// fmt.Println(user1)

	// john := user{
	// 	name:  "John Doe",
	// 	email: "johndoe@go.dev",
	// 	metaInfo: additionalInfo{
	// 		phone:   1234567890,
	// 		address: "123 Main Street",
	// 	},
	// }

	// fmt.Printf("%+v", john)

	newUser := func(name string, age int, role string) user {
		if age <= 0 {
			age = 18
		}

		return user{
			name: name,
			age:  age,
			role: role,
		}
	}

	john := newUser("John Doe", -30, "admin")

	fmt.Print(john)

}

// Create struct type
// Different way to create struct
// Using print to print struct (%+v)
// Updating struct fields
// Accessing struct fields
