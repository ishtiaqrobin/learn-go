package main

import "fmt"

// func multiplyBy(factor int) func(int) int {
// 	return func(x int) int {
// 		return x * factor
// 	}
// }

func makeCounter() func() int {
	count := 0

	inner := func() int {
		count++
		return count
	}

	return inner
}

func main() {
	// double := multiplyBy(2)
	// fmt.Println(double(10))

	next := makeCounter()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
}

// closure

// Closure হলো এমন একটি ফাংশন, যা তার চারপাশের পরিবেশের ভ্যারিয়েবলগুলোকে (Lexical Environment) মেমোরির Heap-এ বন্দি করে নিজের সাথে আজীবনের জন্য বেঁধে রাখে। তাই বাইরের ফাংশনটি মারা গেলেও ভেতরের ফাংশনটি তার হিপের ডাটা ব্যবহার করে বেঁচে থাকে।

// ক্লোজার শুধুমাত্র ভ্যারিয়েবলকে বাঁচিয়েই রাখে না, বরং প্রতিবার কল করার মাঝের সময়ে ভ্যারিয়েবলের মান বা State কেমন ছিল, সেটাও মেমোরিতে নিখুঁতভাবে মনে রাখে।
