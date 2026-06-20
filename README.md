# Go Variables, Constants, and Formatting Examples

## 1. Variables and Constants

```go
package main

import "fmt"

func main() {
	// var name string = "John Doe"
	// role := "admin"

	// var name = "Robin"
	// userName := "Robin"
	// isMarried := true
	// var isMarried bool = true

	// Grouped variables
	// var (
	// 	name string = "Robin"
	// 	age  int    = 23
	// )

	// Multiple variables
	// var position, salary = "Developer", 30000

	const pi = 3.1416

	var point = 5
	point = 4

	fmt.Println(pi)
	fmt.Println(point)
}
```

### Output

```text
3.1416
4
```

---

## 2. Zero Values in Go

When variables are declared without initialization, Go automatically assigns them a **zero value**.

```go
package main

import "fmt"

func main() {
	var age int
	fmt.Println(age) // 0

	var name string
	fmt.Println(name) // ""

	var isAdmin bool
	fmt.Println(isAdmin) // false

	var score float32
	fmt.Println(score) // 0

	var pointer *int
	fmt.Println(pointer) // <nil>
}
```

### Output

```text
0

false
0
<nil>
```

---

## 3. String Formatting with `fmt.Printf`

`fmt.Printf()` allows formatted output using format verbs.

```go
package main

import "fmt"

func main() {
	var name string = "Ishtiaq Robin"
	age := 23
	rating := 4.58585

	fmt.Printf(
		"Hello, I am %s and I am %d years old and my rating is %.1f",
		name,
		age,
		rating,
	)
}
```

### Output

```text
Hello, I am Ishtiaq Robin and I am 23 years old and my rating is 4.6
```

### Common Format Verbs

| Verb   | Description                |
| ------ | -------------------------- |
| `%s`   | String                     |
| `%d`   | Integer                    |
| `%f`   | Float                      |
| `%.1f` | Float with 1 decimal place |
| `%t`   | Boolean                    |
| `%v`   | Default format             |

---

## 4. String Formatting with `fmt.Sprintf`

`fmt.Sprintf()` works like `fmt.Printf()`, but returns the formatted string instead of printing it.

```go
package main

import "fmt"

func main() {
	var name string = "Ishtiaq Robin"
	age := 23
	rating := 4.58585

	formattedString := fmt.Sprintf(
		"Hello, I am %s and I am %d years old and my rating is %.1f",
		name,
		age,
		rating,
	)

	fmt.Println(formattedString)
}
```

## 5. Conditional Statements

```go
package main

import "fmt"

func makeCoffee(kind string, isSugar bool) {
	fmt.Printf("Making %s coffee. \n", kind)
	if isSugar {
		fmt.Println("Adding sugar.")
	}
}

func main() {
	makeCoffee("cappuccino", true)
	makeCoffee("espresso", false)
	makeCoffee("latte", true)
}
```

## 6. Unnamed and Named Return Values

```go
package main

import "fmt"

// unamed return values
// func makeCoffee(kind string) (string, int) {
// 	price := 20
// 	coffee := fmt.Sprintf("%s coffee. \n", kind)
// 	return coffee, price
// }

// named return values
func makeCoffee(kind string) (coffee string, price int) {
	price = 20
	coffee = fmt.Sprintf("%s coffee. \n", kind)

	return
}

func main() {
	myCoffee, myBill := makeCoffee("cappuccino")

	fmt.Printf("I made a %d$ for %s", myBill, myCoffee)
}
```

```go
package main

import "fmt"

func main() {
	// anonymous function
	// makeCoffee := func() {
	// 	fmt.Printf("Making Coffee")
	// }

	// makeCoffee()

	// Immediately Invoked Function Expression - IIFE
	// func() {
	// 	println("Hello World")
	// }()

	// IIFE with parameters
	func(coffeeType string) {
		fmt.Printf("Making %s coffee", coffeeType)
	}("Espresso")
}
```

### Output

```text
Hello, I am Ishtiaq Robin and I am 23 years old and my rating is 4.6
```

---

## Key Takeaways

- Use `var` for explicit variable declarations.
- Use `:=` for short variable declarations inside functions.
- Use `const` for values that should not change.
- Go provides default zero values for uninitialized variables.
- `fmt.Printf()` prints formatted output directly.
- `fmt.Sprintf()` returns a formatted string for later use.
