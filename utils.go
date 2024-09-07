package calendar

import "time"

// DaysInMonth returns the number of days in the given month and year
func DaysInMonth(month, year int) int {
	// Handle February in leap years
	if month == 2 {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 29
		}
		return 28
	}

	// Handle months with 30 or 31 days
	switch month {
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

// DayOfWeek returns the day of the week for the given day, month, and year. (0 is Monday, 6 is Sunday)
func DayOfWeek(day, month, year int) int {
	// Create a time.Time object using the provided day, month, and year
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// Adjust the weekday so Monday is 0
	weekday := int(date.Weekday())
	if weekday == 0 {
		weekday = 6 // Sunday becomes 6
	} else {
		weekday -= 1 // Shift everything else down by 1
	}
	return weekday
}
