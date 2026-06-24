package main

import "fmt"

// type user struct {
// 	name       string
// 	age        int
// 	isLoggedIn bool
// 	greet      func()
// }

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	// user1 := user{
	// 	name:       "Ishtiaq",
	// 	age:        23,
	// 	isLoggedIn: true,
	// }

	// user1.greet = func() {
	// 	println("Hello World")
	// }

	// user1.greet()

	user1 := user{
		name:       "Ishtiaq",
		age:        23,
		isLoggedIn: false,
	}

	user1.greet()

	// eikhane & use kora optional, karon go automatically pointer pass kore
	// user1.login()

	// amra eivabeo pointer use korte pari
	pointerUser1 := &user1
	pointerUser1.login()

	// Alternative way to call a pointer
	// (&user1).login() // গো ব্যাকগ্রাউন্ডে এই রূপান্তরটি নিজে নিজেই করে নেয়!

	fmt.Printf("%+v", user1)
}

// zodi amra * sign use kori, tahole uporer function call er somoy & use na korleo cholbe. karon go automatically pointer pass kore.
func (u *user) login() {
	fmt.Println("Login called")
	u.isLoggedIn = true
}

func (u user) greet() {
	fmt.Println("Hello!", u.name)
}
