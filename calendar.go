package calendar

import (
	"github.com/anotherhadi/markdown"
)

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

type Event struct {
	Name        string
	Description string
	StartDate   Date
	EndDate     Date // If EndDate is not set, the event is considered to be an all-day event
	AllDay      bool
	Settings    map[string]string
}

type Calendar struct {
	Name   string
	Path   string
	Events []Event
	md     *markdown.MarkdownFile
}
