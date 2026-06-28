package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var counter int

func main() {
	for range 1000 {
		wg.Go(increment)
	}

	wg.Wait()

	fmt.Println("Counter value is", counter)
}

func increment() {
	mu.Lock()         // ১. দরজা লক করে ভেতরে ঢোকা
	defer mu.Unlock() // ৩. ফাংশন শেষ হওয়ার মুহূর্তে দরজা খুলে দেওয়া
	counter++         // ২. ক্রিপ্টোগ্রাফিক বা ক্রিটিক্যাল জোনের কাজ (Read-Modify-Write)
}
