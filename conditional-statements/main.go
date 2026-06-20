package main

import "fmt"

func makeCoffee(kind string, isSugar bool) {
	fmt.Printf("Making %s coffee. \n", kind)
	if isSugar {
		fmt.Println("Adding sugar.")
	}
}

func main() {
	makeCoffee("cappuccino", true)
	makeCoffee("espresso", false)
	makeCoffee("latte", true)
}
