package main

import "fmt"

func main() {
	fmt.Println("=== RUNNING MULTI-FILE ARCHITECTURE ===")

	// ১. bKash টেস্ট (পয়েন্টার ফ্লো - মেমোরি অ্যাড্রেস আপডেট হবে)
	myBkash := Bkash{WalletNumber: "01700000000", Balance: 1000.00}
	bkashService := NewPaymentService(&myBkash) // ঠিকানা পাঠানো হলো
	bkashService.Checkout(200.00)
	fmt.Printf("[Main] Main memory bKash balance: %.2f tk\n\n", myBkash.Balance)

	// ২. Mock টেস্ট (ভ্যালু ফ্লো - মেমোরি অ্যাড্রেস আপডেট হবে না)
	myMock := MockPayment{IsTested: false}
	mockService := NewPaymentService(myMock) // সরাসরি ভ্যালু পাঠানো হলো
	mockService.Checkout(10.00)
	fmt.Printf("[Main] Main memory Mock IsTested - value: %t\n", myMock.IsTested)
}

// * For run the file
// go run main.go interfaces.go methods.go mock.go

//* or
// go run .
