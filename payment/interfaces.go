package main

// =========================================================================
// ১. INTERFACE (চুক্তিনামা)
// =========================================================================
// কে কী চায়: এটি শুধু 'Pay' মেথডের নিয়মটি ধরে রাখে।
type PaymentProcessor interface {
	Pay(amount float64)
}

// =========================================================================
// ২. MASTER SERVICE BOX (পেমেন্ট প্রসেস করার ইঞ্জিন)
// =========================================================================
type PaymentService struct {
	method PaymentProcessor // ইন্টারফেসের খালি স্লট
}

// CONSTRUCTOR: বাইরে থেকে পেমেন্ট মেথড রিসিভ করে সার্ভিসের ভেতর ইনজেক্ট করে
func NewPaymentService(mp PaymentProcessor) *PaymentService {
	return &PaymentService{
		method: mp,
	}
}

// CHECKOUT: ইন্টারফেসে থাকা নির্দিষ্ট মেথডের Pay() কে ট্রিগার করে
func (ps PaymentService) Checkout(amount float64) {
	ps.method.Pay(amount) // এখান থেকে কোড লাফ দিয়ে নির্দিষ্ট ফাইলের Pay() মেথডে চলে যায়
}
