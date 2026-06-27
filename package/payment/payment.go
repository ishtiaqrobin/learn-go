package payment

import "fmt"

type paymentMethod interface {
	Pay(amount float64)
}

type bKash struct {
	apiKey string
}

type Rocket struct {
	apiKey string
}

func (bk *bKash) Pay(amount float64) {
	// actual payment logic
	fmt.Printf("Paying %.2f tk with bKash \n", amount)
}

func (r *Rocket) Pay(amount float64) {
	// actual payment logic
	fmt.Printf("Paying %.2f tk with Rocket", amount)
}

type PaymentService struct {
	method paymentMethod
}

func NewPaymentService(method paymentMethod) *PaymentService {
	return &PaymentService{
		method: method,
	}
}

func NewRocket(apiKey string) *Rocket {
	return &Rocket{
		apiKey: apiKey,
	}
}

func (ps PaymentService) Checkout() {
	ps.method.Pay(10.00)
}
