package api

import (
	"fmt"
	"time"
)

// Day represents the day of week.
type Day int

// Days for setting alarms.
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (d Day) String() string {
	switch d {
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	case Sunday:
		return "Sunday"
	default:
		return fmt.Sprintf("Day(%d)", d)
	}
}

// FromWeekday converts the provided weekday to Day format.
func FromWeekday(d time.Weekday) Day {
	switch d {
	case time.Sunday:
		return Sunday
	default:
		return Day(d - 1)
	}
}
