package calendar

import (
	"strconv"

	"github.com/anotherhadi/markdown"
)

func (c *Calendar) AddEvent(event Event) {
	c.Events = append(c.Events, event)
	if event.EndDate == (Date{}) {
		c.Events[len(c.Events)-1].AllDay = true
	}

	c.md.AddSection("## " + event.Name)

	section := c.md.GetSection(markdown.H2, event.Name)
	for key, value := range event.Settings {
		section.AddLine("- " + key + ": " + value)
	}
	if event.StartDate != (Date{}) {
		section.AddLine("- start_date: " + strconv.Itoa(event.StartDate.Year) + "-" + strconv.Itoa(event.StartDate.Month) + "-" + strconv.Itoa(event.StartDate.Day) + " " + strconv.Itoa(event.StartDate.Hour) + ":" + strconv.Itoa(event.StartDate.Minute))
	}
	if event.EndDate != (Date{}) {
		section.AddLine("- end_date: " + strconv.Itoa(event.EndDate.Year) + "-" + strconv.Itoa(event.EndDate.Month) + "-" + strconv.Itoa(event.EndDate.Day) + " " + strconv.Itoa(event.EndDate.Hour) + ":" + strconv.Itoa(event.EndDate.Minute))
	}
	section.AddLine(event.Description)
}

func (c *Calendar) Write(str ...string) (err error) {
	if len(str) == 0 {
		return c.md.Write()
	}
	return c.md.Write(str[0])
}
