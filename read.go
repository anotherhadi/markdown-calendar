package calendar

import (
	"strconv"
	"strings"

	"github.com/anotherhadi/markdown"
)

func Read(path string) (Calendar, error) {
	c := Calendar{Path: path}
	md := markdown.New(path)
	err := md.Read()
	if err != nil {
		return Calendar{}, err
	}
	c.md = &md

	c.Name = md.Title
	c.EventColor = c.md.GetFrontMatter("event_color", "#A594FD").(string)

	// Load events
	for _, section := range md.Sections {
		if section.SectionType != markdown.H2 {
			continue
		}
		event := Event{}
		event.Name = section.Text
		event.Description = ""
		event.Settings = make(map[string]string)
		for _, line := range section.Lines {
			if line.LineType == markdown.List {
				linestr := strings.TrimSpace(line.Text)
				linestr = strings.TrimPrefix(linestr, "- ")
				if strings.Contains(linestr, ": ") {
					key := strings.TrimSpace(strings.Split(linestr, ": ")[0])
					value := strings.TrimSpace(strings.Split(linestr, ": ")[1])
					event.Settings[key] = value
				}
				continue
			}
			if line.Text == "" {
				continue
			}
			event.Description += line.Text + "\n"
		}
		c.Events = append(c.Events, event)
	}

	// Parse dates
	for i, event := range c.Events {
		if start, ok := event.Settings["start_date"]; ok {
			date := strings.Split(start, " ")[0]
			if len(strings.Split(date, "-")) == 3 {
				c.Events[i].StartDate.Year, _ = strconv.Atoi(strings.Split(date, "-")[0])
				c.Events[i].StartDate.Month, _ = strconv.Atoi(strings.Split(date, "-")[1])
				c.Events[i].StartDate.Day, _ = strconv.Atoi(strings.Split(date, "-")[2])
			}
			hour := strings.Split(start, " ")
			if len(hour) > 1 {
				c.Events[i].StartDate.Hour, _ = strconv.Atoi(strings.Split(hour[1], ":")[0])
				c.Events[i].StartDate.Minute, _ = strconv.Atoi(strings.Split(hour[1], ":")[1])
			}
		}
		if start, ok := event.Settings["end_date"]; ok {
			date := strings.Split(start, " ")[0]
			if len(strings.Split(date, "-")) == 3 {
				c.Events[i].EndDate.Year, _ = strconv.Atoi(strings.Split(date, "-")[0])
				c.Events[i].EndDate.Month, _ = strconv.Atoi(strings.Split(date, "-")[1])
				c.Events[i].EndDate.Day, _ = strconv.Atoi(strings.Split(date, "-")[2])
			}
			hour := strings.Split(start, " ")
			if len(hour) > 1 {
				c.Events[i].EndDate.Hour, _ = strconv.Atoi(strings.Split(hour[1], ":")[0])
				c.Events[i].EndDate.Minute, _ = strconv.Atoi(strings.Split(hour[1], ":")[1])
			}
		} else {
			c.Events[i].AllDay = true
		}
	}

	return c, nil
}

func (c *Calendar) GetEventsByDate(year, month, day int) []Event {
	var events []Event
	for _, event := range c.Events {
		if event.StartDate.Year == year && event.StartDate.Month == month && event.StartDate.Day == day {
			events = append(events, event)
		}
	}
	return events
}

func (c *Calendar) GetEventsByMonth(year, month int) []Event {
	var events []Event
	for _, event := range c.Events {
		if event.StartDate.Year == year && event.StartDate.Month == month {
			events = append(events, event)
		}
	}
	return events
}

func (c *Calendar) GetEventsByYear(year int) []Event {
	var events []Event
	for _, event := range c.Events {
		if event.StartDate.Year == year {
			events = append(events, event)
		}
	}
	return events
}

func (c *Calendar) GetEventsByName(name string) []Event {
	var events []Event
	for _, event := range c.Events {
		if event.Name == name {
			events = append(events, event)
		}
	}
	return events
}
