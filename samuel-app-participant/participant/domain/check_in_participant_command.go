package domain

type CheckInParticipant struct {
	EventId string `json:"eventId" validate:"required,uuid4"`
	Address CheckInParticipantAddress `json:"address" validate:"required"`
}

type CheckInParticipantAddress struct {
	Street string `json:"street" validate:"required"`
	Number int `json:"number" validate:"gt=0"`
	ZipCode string `json:"zipCode" validate:"len=5"`
	City    string `json:"city" validate:"required"`
}
