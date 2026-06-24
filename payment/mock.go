package main

import "fmt"

// =========================================================================
// MockPayment (নকল টেস্টিং গেটওয়ে)
// =========================================================================
type MockPayment struct {
	IsTested bool
}

// Value Receiver: পয়েন্টার না থাকায় এটি আসল মেমোরি আপডেট না করে শুধু কপিতে কাজ করে
func (m MockPayment) Pay(amount float64) {
	m.IsTested = true
	fmt.Printf("[Mock] Payment test is successful. (Copied IsTested = %t)\n", m.IsTested)
}
