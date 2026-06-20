package main

import "fmt"

// unamed return values
// func makeCoffee(kind string) (string, int) {
// 	price := 20
// 	coffee := fmt.Sprintf("%s coffee. \n", kind)
// 	return coffee, price
// }

// named return values
func makeCoffee(kind string) (coffee string, price int) {
	price = 20
	coffee = fmt.Sprintf("%s coffee. \n", kind)

	return
}

func main() {
	myCoffee, myBill := makeCoffee("cappuccino")

	fmt.Printf("I made a %d$ for %s", myBill, myCoffee)
}
