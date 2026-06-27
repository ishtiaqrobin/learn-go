package main // executable package

import (
	"fmt"
	"learngo/payment"
	"learngo/test"
)

func init() {
	fmt.Println("Initializing db...")
}

func print() {
	fmt.Println("I am in test function")
}

func main() {
	// bkash := bKash{apiKey: "1234abc"}
	// paymentServiceBkash := NewPaymentService(&bkash)
	// paymentServiceBkash.checkout()

	// rocket := NewRocket("1234abc")
	// paymentServiceRocket := NewPaymentService(rocket)
	// paymentServiceRocket.checkout()

	mockPm := test.MockPaymentMethod{}

	paymentServiceMock := payment.NewPaymentService(mockPm)
	paymentServiceMock.Checkout()

	// Print with default helper functions
	// color.Cyan("Prints text in cyan.")

	fmt.Println("I am in main function")

	print()
}

// module = bunch of package
