package main

import "fmt"

// =========================================================================
// কোম্পানি ১: bKash (পয়েন্টার ব্যাটারি)
// =========================================================================
// কে কী চায়: bKash অবজেক্টটি WalletNumber আর Balance চায়।
type Bkash struct {
	WalletNumber string
	Balance      float64
}

// Pointer Receiver: পয়েন্টার থাকায় এটি আসল মেমোরি আপডেট করে
func (bk *Bkash) Pay(amount float64) {
	bk.Balance -= amount
	fmt.Printf("[bKash] %.2f tk Payment is Successful. Rest of your balance: %.2f tk\n", amount, bk.Balance)
}

// =========================================================================
// কোম্পানি ২: Rocket (ফ্যাক্টরি-মেড ব্যাটারি)
// =========================================================================
type Rocket struct {
	ApiKey string
}

// ফ্যাক্টরি ফাংশন: Rocket অবজেক্ট তৈরি করে পয়েন্টার রিটার্ন করে
func NewRocket(apiKey string) *Rocket {
	return &Rocket{ApiKey: apiKey}
}

// Pointer Receiver
func (r *Rocket) Pay(amount float64) {
	fmt.Printf("[Rocket] %.2f tk Payment is Successful. (API Key: %s)\n", amount, r.ApiKey)
}
