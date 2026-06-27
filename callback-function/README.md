# What I'm Learning

This repo is a hands-on workspace for learning Go. Each folder in `Mission-9/` covers a different concept.

## Topics Covered

### 1. Callback Functions
(`callback-function/`)

What I learned:

- Go treats **functions as first-class citizens** — they can be assigned to variables, passed as arguments, and returned from other functions.
- A **callback** is a function that is passed into another function to be invoked later.
- Example: `calculator(a, b, operation)` accepts any function with the signature `func(int, int) int`, so I can plug in `add`, `multiply`, or any custom logic without changing `calculator` itself.

Key pattern:
```go
func calculator(a int, b int, operation func(x int, y int) int) int {
    return operation(a, b)
}

add      := func(n1, n2 int) int { return n1 + n2 }
multiply := func(n1, n2 int) int { return n1 * n2 }

fmt.Println(calculator(10, 20, add))      // 30
fmt.Println(calculator(10, 20, multiply)) // 200
```

## How to run

```bash
cd E:\Level2-B6\Mission-9\callback-function
go run main.go
```

Module name: `learngo` · Go version: 1.26.4
