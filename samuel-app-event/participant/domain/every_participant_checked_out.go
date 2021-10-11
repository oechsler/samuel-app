package domain

type EveryParticipantCheckedOut struct {
	EventId string `json:"eventId"`
}

func NewEveryParticipantCheckedOut(eventId string) *EveryParticipantCheckedOut {
	return &EveryParticipantCheckedOut{EventId: eventId}
}