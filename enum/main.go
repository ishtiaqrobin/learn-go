package main

import "fmt"

type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// 1. OfficeStatus is defined as a custom type.
type OfficeStatus string

const (
	WorkingDay OfficeStatus = "open"
	HalfDay    OfficeStatus = "half_day"
	OffDay     OfficeStatus = "closed"
)

// 2. Return type change string to OfficeStatus
func getWorkDayStatus(day Weekday) OfficeStatus {
	switch day {
	case Sunday, Monday, Tuesday, Wednesday:
		return WorkingDay // Direct enum constant return
	case Thursday:
		return HalfDay
	case Friday, Saturday:
		return OffDay
	default:
		// If the day is not a valid weekday, return an invalid status
		return "invalid_status"
	}
}

func main() {
	status := getWorkDayStatus(Thursday)

	fmt.Println("Office Status:", status) // output: half_day

	// You can also compare it directly with the enum if you want.
	if status == HalfDay {
		fmt.Println("It's a half day.")
	}
}
