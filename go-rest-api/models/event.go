package models

import "time"

type Event struct {
	ID          int
	// tag to make binding to be required in gin
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
