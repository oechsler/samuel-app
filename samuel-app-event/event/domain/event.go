package domain

import (
	"time"
)

type Event struct {
	Id string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	StartedAt time.Time `bson:"startedAt" json:"startedAt"`
	CompletedAt time.Time `bson:"completedAt" json:"completedAt"`
}

func NewEvent(id string, name string) *Event {
	return &Event{Id: id, Name: name}
}

func (e Event) Started(startedAt string) *Event {
	if !e.StartedAt.Equal(time.Time{}) {
		return &e
	}
	e.StartedAt, _ = time.Parse(time.RFC3339, startedAt)

	return &e
}

func (e Event) Completed(completedAt string) *Event {
	if !e.CompletedAt.Equal(time.Time{}) {
		return &e
	}
	e.CompletedAt, _ = time.Parse(time.RFC3339, completedAt)

	return &e
}

func (e Event) IsActive() bool {
	return !e.StartedAt.Equal(time.Time{}) && e.CompletedAt.Equal(time.Time{})
}