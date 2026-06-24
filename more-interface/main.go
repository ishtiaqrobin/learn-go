package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type bKash struct {
	apiKey string
}

type Rocket struct {
	apiKey string
}

func (bk *bKash) pay(amount float64) {
	// actual payment logic
	fmt.Printf("Paying %.2f tk with bKash \n", amount)
}

func (r *Rocket) pay(amount float64) {
	// actual payment logic
	fmt.Printf("Paying %.2f tk with Rocket", amount)
}

type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *PaymentService {
	return &PaymentService{
		method: method,
	}
}

func NewRocket(apiKey string) *Rocket {
	return &Rocket{
		apiKey: apiKey,
	}
}

func (ps PaymentService) checkout() {
	ps.method.pay(10.00)
}

type MockPaymentMethod struct{}

func (mockPM MockPaymentMethod) pay(amount float64) {
	// test logic
	fmt.Println("Testing payment method")
}

func main() {
	// bkash := bKash{apiKey: "1234abc"}
	// paymentServiceBkash := NewPaymentService(&bkash)
	// paymentServiceBkash.checkout()

	// rocket := NewRocket("1234abc")
	// paymentServiceRocket := NewPaymentService(rocket)
	// paymentServiceRocket.checkout()

	mockPm := MockPaymentMethod{}
	paymentServiceMock := NewPaymentService(mockPm)
	paymentServiceMock.checkout()
}
