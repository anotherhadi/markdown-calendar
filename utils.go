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

func IncrementYear(day, month, year *int, inc int) {
	*year += inc
	if *year < 1 {
		*year = 1
	}
	if *year > 9999 {
		*year = 9999
	}
	if *day > DaysInMonth(*month, *year) {
		*day = DaysInMonth(*month, *year)
	}
}

func IncrementMonth(day, month, year *int, inc int) {
	*month += inc
	if *month < 1 {
		*month = 12
		IncrementYear(day, month, year, -1)
	}
	if *month > 12 {
		*month = 1
		IncrementYear(day, month, year, 1)
	}
	if *day > DaysInMonth(*month, *year) {
		*day = DaysInMonth(*month, *year)
	}
}

func IncrementDay(day, month, year *int, inc int) {
	tmp := *day
	tmp += inc
	if tmp < 1 {
		IncrementMonth(day, month, year, -1)
		*day = DaysInMonth(*month, *year) + inc + *day
	} else if tmp > DaysInMonth(*month, *year) {
		*day = *day + inc - DaysInMonth(*month, *year)
		IncrementMonth(day, month, year, 1)
	} else {
		*day = tmp
	}
}
