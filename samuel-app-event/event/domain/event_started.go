package domain

import (
	"time"
)

type EventStarted struct {
	Id string `json:"id"`
	StartedAt string `json:"startedAt"`
}

func NewEventStarted(id string) *EventStarted {
	startedAt := time.Now().UTC().Format(time.RFC3339)

	return &EventStarted{Id: id, StartedAt: startedAt}
}
