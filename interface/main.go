// package main

// import "fmt"

// // ts => interface for data shape
// // go => interface behavior contract

// // ১. আমরা একটি ইন্টারফেস বা চুক্তিনামা বানালাম
// // যে-ই 'speak()' মেথডটি ধারণ করবে, সে-ই একটি 'Speaker'
// type Speaker interface {
// 	speak()
// }

// type Dog struct{}

// // ২. Dog স্ট্রাকটের জন্য speak() মেথডটি লিখলাম
// // (পয়েন্টার জটিলতা এড়াতে আপাতত ভ্যালু রিসিভার ব্যবহার করলাম)
// func (d Dog) speak() {
// 	fmt.Println("Woof!")
// }

// type Cat struct{}

// // ৩. ইন্টারফেসের পাওয়ার দেখানোর জন্য আরেকটি Cat স্ট্রাকট বানালাম
// func (c Cat) speak() {
// 	fmt.Println("Meow!")
// }

// type Human struct {
// 	name string // name is encapsulated because it is private and it smaller letter start
// }

// func (h Human) speak() {
// 	fmt.Println("Hello! My name is :", h.name)
// }

// func makeSound(s Speaker) {
// 	s.speak()
// }

// func main() {
// 	myDog := Dog{}
// 	myCat := Cat{}
// 	human := Human{name: "Ishtiaq Robin"}

// 	// Dog এবং Cat দুজনেই Speaker ইন্টারফেসের নিয়ম মেনে চলেছে,
// 	// তাই দুজনেই makeSound ফাংশনে পাস হতে পারবে।
// 	makeSound(myDog) // আউটপুট: Woof!
// 	makeSound(myCat) // আউটপুট: Meow!
// 	makeSound(human) // আউটপুট: Hello!
// }

// // OOP
// // Abstraction, Encapsulation, Inheritance, Polymorphism

package main

import "fmt"

// =========================================================================
// ১. ABSTRACTION (অ্যাবস্ট্রাকশন) - চুক্তিনামা বা রিমোট কন্ট্রোল
// =========================================================================
// এখানে আমরা একটি পেমেন্ট গেটওয়ের 'নকশা' তৈরি করলাম।
// সিস্টেমের ভেতরের জটিল কোড কেমন হবে তা আমাদের জানার দরকার নেই,
// শুধু এই বাটনটি (Pay) থাকলেই চলবে।
type PaymentGateway interface {
	Pay(amount float64) // যে-ই এই মেথডটি লিখবে, সে-ই পেমেন্ট করতে পারবে।
}

// =========================================================================
// ২. ENCAPSULATION (এনক্যাপসুলেশন) - ডাটা বা প্রোপার্টি লক করা
// =========================================================================
type Bkash struct {
	// ফিল্ডের নাম ছোট হাতের অক্ষরে (walletNumber) দেওয়ায় এটি এই প্যাকেজের বাইরে Private।
	// বাইরের কোনো কোড হুট করে এসে সরাসরি এই নাম্বারটি ওলটপালট করতে পারবে না।
	walletNumber string
}

// এটি একটি কনস্ট্রাকটর ফাংশন, যা সুরক্ষিতভাবে Bkash অবজেক্ট তৈরি করতে সাহায্য করে।
func NewBkash(number string) Bkash {
	return Bkash{walletNumber: number}
}

// =========================================================================
// ৩. POLYMORPHISM (পলিমরফিজম) - একই মেথডের ভিন্ন ভিন্ন রূপ বা আচরণ
// =========================================================================

// Bkash-এর নিজস্ব স্টাইলে পেমেন্ট করার নিয়ম
func (b Bkash) Pay(amount float64) {
	fmt.Printf("[bKash] %0.2f টাকা সফলভাবে পেমেন্ট হয়েছে (নম্বর: %s)\n", amount, b.walletNumber)
}

// আরেকটি পেমেন্ট মেথড: Rocket
type Rocket struct {
	accountID string
}

// Rocket-এর নিজস্ব স্টাইলে পেমেন্ট করার নিয়ম (আলাদা আচরণ বা রূপ)
func (r Rocket) Pay(amount float64) {
	fmt.Printf("[Rocket] %0.2f টাকা চার্জসহ প্রসেস করা হয়েছে (আইডি: %s)\n", amount, r.accountID)
}

// 💡 পলিমরফিজমের আসল ম্যাজিক এই ফাংশনটি!
// এটি কোনো নির্দিষ্ট Bkash বা Rocket চেনে না। সে শুধু চেনে 'PaymentGateway' ইন্টারফেস।
// আমরা একে যে পেমেন্ট মেথডই দেব, সে সেটার রূপ ধারণ করে কাজ করে দেবে।
func ProcessOrder(gateway PaymentGateway, bill float64) {
	gateway.Pay(bill) // এখানে রানটাইমে নির্ধারিত হবে কোন Pay() কল হচ্ছে
}

// =========================================================================
// ৪. COMPOSITION (কম্পোজিশন) - ইনহেরিটেন্সের বিকল্প (ছোট অংশ জোড়া লাগানো)
// =========================================================================
type User struct {
	Name string
}

type PremiumUser struct {
	User               // 💡 Struct Embedding! User-এর সব বৈশিষ্ট্য (Name) এখন PremiumUser পেয়ে গেল।
	DiscountPercentage float64
}

// =========================================================================
// MAIN FUNCTION - রান করে টেস্ট করার জায়গা
// =========================================================================
func main() {
	fmt.Println("--- OOP Concepts in Go ---")

	// ক. এনক্যাপসুলেটেড অবজেক্ট তৈরি
	bkashPayment := NewBkash("01711223344")
	rocketPayment := Rocket{accountID: "998877"}

	// খ. কম্পোজিশন বা এমবেডিং টেস্ট
	// PremiumUser তৈরি করছি, দেখবে সে সরাসরি 'Name' অ্যাক্সেস করতে পারছে
	vipCustomer := PremiumUser{
		User:               User{Name: "Ishtiaq Robin"},
		DiscountPercentage: 10.0,
	}

	fmt.Printf("কাস্টমার: %s (প্রিমিয়াম মেম্বার)\n\n", vipCustomer.Name)

	// গ. পলিমরফিজম টেস্ট (একই ফাংশনে ভিন্ন ভিন্ন মেথড পাস করা)
	fmt.Println("১ম অর্ডারের পেমেন্ট প্রসেস হচ্ছে:")
	ProcessOrder(bkashPayment, 500.00) // Bkash-এর Pay() কল হবে

	fmt.Println("\n২য় অর্ডারের পেমেন্ট প্রসেস হচ্ছে:")
	ProcessOrder(rocketPayment, 1200.50) // Rocket-এর Pay() কল হবে
}
