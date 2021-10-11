package domain

import uuid "github.com/satori/go.uuid"

type EventPrepared struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func NewEventPrepared(name string) *EventPrepared {
	return &EventPrepared{Id: uuid.NewV4().String(), Name: name}
}