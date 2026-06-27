package test

import "fmt"

func init() {
	fmt.Println("I am from test init function")
}

type MockPaymentMethod struct{}

func (mockPM MockPaymentMethod) Pay(amount float64) {
	// test logic
	fmt.Println("Testing payment method")
}
