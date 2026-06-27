package main

import "encoding/json"

type person struct {
	Name string
	Age  int
	City string
}

func main() {
	p := person{
		Name: "Ishtiaq",
		Age:  23,
		City: "Dhaka",
	}

	json.Marshal(p)
}
