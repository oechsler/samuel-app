package domain

import (
	"time"
)

type EventCompleted struct {
	Id string `json:"id"`
	CompletedAt string `json:"completedAt"`
}

func NewEventCompleted(id string) *EventCompleted {
	completedAt := time.Now().UTC().Format(time.RFC3339)

	return &EventCompleted{Id: id, CompletedAt: completedAt}
}
