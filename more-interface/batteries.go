package main

import "fmt"

// =========================================================================
// কোম্পানি ১: bKash ব্যাটারি (The Pointer Battery)
// =========================================================================
// কে কী চায়: বিকাশ অবজেক্টটি তৈরি হওয়ার জন্য শুধুমাত্র একটি 'apiKey' চায়।
type bKash struct {
	apiKey string
}

// ওয়ার্কফ্লো: বিকাশের মেথডটি 'Pointer Receiver (*bKash)' দিয়ে লেখা।
// এর মানে হলো, ঘড়ির স্লটে এই ব্যাটারিটা দিতে হলে সরাসরি দিলে হবে না,
// এর মেমোরি ঠিকানা বা অ্যাড্রেস (&) দিতে হবে, যেন ঘড়ি সরাসরি আসল ব্যাটারি থেকে পাওয়ার নিতে পারে।
func (bk *bKash) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with bKash (Using API Key: %s)\n", amount, bk.apiKey)
}

// =========================================================================
// কোম্পানি ২: Rocket ব্যাটারি (The Factory-Made Battery)
// =========================================================================
// কে কী চায়: রকেট অবজেক্টটিও একটি 'apiKey' চায়।
type Rocket struct {
	apiKey string
}

// ওয়ার্কফ্লো: রকেট কোম্পানি একটা হেল্পার ফাংশন (NewRocket) বানিয়ে দিয়েছে।
// তুমি একে শুধু apiKey দেবে, সে নিজেই রকেট ব্যাটারি তৈরি করে তার অ্যাড্রেস (*Rocket) রিটার্ন করবে।
func NewRocket(apiKey string) *Rocket {
	return &Rocket{
		apiKey: apiKey,
	}
}

// রকেটের ২ ইঞ্চির মেথড যা ঘড়ির স্লটের সাথে ম্যাচ করবে।
func (r *Rocket) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Rocket (Using API Key: %s)\n", amount, r.apiKey)
}

// =========================================================================
// কোম্পানি ৩: MockPaymentMethod (The Paper Battery)
// =========================================================================
// কে কী নেয়/চায়: এর ভেতরে কোনো ডাটা বা এপিআই কি-র দরকার নেই। এটি সম্পূর্ণ খালি।
type MockPaymentMethod struct{}

// ওয়ার্কফ্লো: এটি একটি 'Value Receiver' (কোনো স্টার * নাই)।
// এর মানে হলো এটি এতই হালকা যে এর কোনো মেমোরি ঠিকানার দরকার নেই,
// সরাসরি এই কাগজের ব্যাটারিটা ঘড়ির স্লটে ছুঁড়ে দিলেই সে ফিট হয়ে যায়।
func (mockPM MockPaymentMethod) pay(amount float64) {
	// আসল পেমেন্ট না করে এটি জাস্ট টেস্ট সাকসেস মেসেজ দেখায়
	fmt.Println("Testing mode active: Clock is ticking perfectly with Mock Battery!")
}
