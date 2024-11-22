package models

import "time"

type Event struct {
	ID int
	Title string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

func (s *Event) Save()  {
	// later: dd it to a database
	events = append(events, *s)
}

func GetAllEvents() []Event {
	return events
}