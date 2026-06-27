package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"personName"` // json tag/struct tag
	Age  int    `json:"-"`          // use - to ignore
	City string `json:"city"`       // use json tag
}

func main() {
	p := person{
		Name: "Ishtiaq",
		Age:  23,
		City: "Dhaka",
	}

	rawJson, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(rawJson))
}
