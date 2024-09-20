package calendar

import (
	"errors"
	"time"

	purple "github.com/anotherhadi/purple-apps"
)

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

func Today() (day, month, year int) {
	t := time.Now()
	return t.Day(), int(t.Month()), t.Year()
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

func (d Date) IsBefore(date Date) bool {
	if d.Year < date.Year {
		return true
	}
	if d.Month < date.Month {
		return true
	}
	if d.Day < date.Day {
		return true
	}
	if d.Hour < date.Hour {
		return true
	}
	if d.Minute < date.Minute {
		return true
	}
	return false
}

func (d Date) IsAfter(date Date) bool {
	return !d.IsBefore(date) &&
		(d.Day != date.Day || d.Month != date.Month || d.Year != date.Year || d.Hour != date.Hour || d.Minute != date.Minute)
}

func (e Event) IsPast() bool {
	day, month, year := Today()
	hour, minute := time.Now().Hour(), time.Now().Minute()
	today := Date{Day: day, Month: month, Year: year, Hour: hour, Minute: minute}
	// if e.EndDate is not set, take e.StartDate
	if e.EndDate == (Date{}) {
		return e.StartDate.IsBefore(today)
	} else {
		return e.EndDate.IsBefore(today)
	}
}

func MergeCalendars(calendars []Calendar) Calendar {
	merged := Calendar{}
	for _, cal := range calendars {
		for _, event := range cal.Events {
			event.CalendarName = cal.Name
			merged.Events = append(merged.Events, event)
		}
	}
	return merged
}

func GetCalendarsNames(calendars []Calendar) []string {
	names := make([]string, len(calendars))
	for i, cal := range calendars {
		names[i] = cal.Name
	}
	return names
}

func (c Calendar) GetColor(defaultColor string) string {
	if c.EventColor != "" {
		return c.EventColor
	}
	return defaultColor
}

func GetCalendarByName(calendars []Calendar, name string) (*Calendar, error) {
	for i, cal := range calendars {
		if cal.Name == name {
			return &calendars[i], nil
		}
	}
	return nil, errors.New("Calendar not found")
}

func GetPurpleCalendars() []*Calendar {
	calendars := []*Calendar{}
	for _, p := range purple.Config.Calendar.Paths {
		cal, err := Read(p)
		if err != nil {
			continue
		}
		calendars = append(calendars, &cal)
	}
	return calendars
}
