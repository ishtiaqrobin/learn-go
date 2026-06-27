package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func getWorkDayStatus(day Weekday) string {
	switch day {
	case Sunday, Monday, Tuesday, Wednesday:
		return "Working day"
	case Thursday:
		return "Half day"
	case Friday, Saturday:
		return "Off day"
	default:
		return "Invalid day"
	}
}

type OfficeStatus string

const (
	WorkingDay OfficeStatus = "open"
	HalfDay    OfficeStatus = "half_day"
	OffDay     OfficeStatus = "closed"
)

func main() {
	fmt.Println(getWorkDayStatus(Thursday))
}
