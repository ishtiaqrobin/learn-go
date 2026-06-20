package main

func main() {
	day := "mon"

	// if day == "mon" {
	// 	fmt.Println("Yes, it is Monday. Sleep more.")
	// } else if day == "tue" {
	// 	fmt.Println("No, it is not Monday. Work hard.")
	// } else {
	// 	fmt.Println("No, it is not Monday. Work hard.")
	// }

	// Tagged switch
	// switch day {
	// case "mon":
	// 	println("Yes, it is Monday. Sleep more.")
	// case "tue":
	// 	println("No, it is not Monday. Work hard.")
	// default:
	// 	println("Another day")
	// }

	// Normal switch
	switch {
	case day == "mon":
		println("Yes, it is Monday. Sleep more.")
	case day == "tue":
		println("No, it is not Monday. Work hard.")
	default:
		println("Another day")
	}
}
