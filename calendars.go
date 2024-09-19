package calendar

func GetEventsByDate(year, month, day int, calendars []Calendar) []Event {
	n := []Event{}
	for _, cal := range calendars {
		n = append(n, cal.GetEventsByDate(year, month, day)...)
	}
	return n
}

func GetNumberOfEventsByDate(year, month, day int, calendars []Calendar) int {
	n := 0
	for _, cal := range calendars {
		n += len(cal.GetEventsByDate(year, month, day))
	}
	return n
}
func GetNumberOfEventsInMonth(year, month int, calendars []Calendar) int {
	n := 0
	for _, cal := range calendars {
		n += len(cal.GetEventsByMonth(year, month))
	}
	return n
}

func GetNumberOfEventsInYear(year int, calendars []Calendar) int {
	n := 0
	for _, cal := range calendars {
		n += len(cal.GetEventsByYear(year))
	}
	return n
}
