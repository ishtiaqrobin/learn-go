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
	// var p person
	// p = person{
	// 	Name: "Ishtiaq",
	// 	Age:  23,
	// 	City: "Dhaka",
	// }

	// rawJson, err := json.Marshal(p)

	// if err != nil {
	// 	fmt.Println("Error", err)
	// }

	// fmt.Println(string(rawJson))

	var p2 person
	jsonText := `{"personName":"Ishtiaq","city":"Dhaka"}`

	error := json.Unmarshal([]byte(jsonText), &p2)
	if error != nil {
		fmt.Println("Error", error)
	}

	fmt.Println(p2.Name)
	fmt.Println(p2.Age)
	fmt.Println(p2.City)

}
