package main

import (
	calendar "github.com/anotherhadi/markdown-calendar"
)

func main() {
	c, _ := calendar.Read("./test/test.md")
	c.AddEvent(calendar.Event{
		Name:        "TestEvent",
		Description: "This is a test event",
		Settings:    map[string]string{"asetting": "woops"},
		StartDate:   calendar.Date{Year: 2020, Month: 1, Day: 1, Hour: 12, Minute: 0},
	})
	err := c.Write("./test/test2.md")
	if err != nil {
		panic(err)
	}

}
