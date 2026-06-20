package main

import "fmt"

func ageStatus(age int) {

}

func main() {
	// age := 20

	// isAdult := age >= 18

	// if isAdult {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are not an adult")
	// }

	if score := 85; score >= 80 {
		fmt.Println("You got an a Golden Medal and your score is: ", score)
	} else if score >= 70 {
		fmt.Println("You got a Silver Medal and your score is: ", score)
	} else {
		fmt.Println("You got a Participant Certificate and your score is: ", score)
	}

	// Example:
	// if err := saveToDB(); err != nil {
	// 	fmt.Println("Error saving to DB: ", err)
	// }
}
