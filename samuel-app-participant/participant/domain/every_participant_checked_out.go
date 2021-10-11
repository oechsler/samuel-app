package domain

type EveryParticipantCheckedOut struct {
	EventId string `json:"eventId"`
}

func NewCheckedOutEveryone(eventId string) *EveryParticipantCheckedOut {
	return &EveryParticipantCheckedOut{EventId: eventId}
}