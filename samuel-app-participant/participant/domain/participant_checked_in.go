package domain

import uuid "github.com/satori/go.uuid"

type ParticipantCheckedIn struct {
	Id string `json:"id"`
	EventId string `json:"eventId"`
	Address *Address `json:"address"`
}

func NewParticipantCheckedIn(eventId string, address *Address) *ParticipantCheckedIn {
	return &ParticipantCheckedIn{Id: uuid.NewV4().String(), EventId: eventId, Address: address}
}
