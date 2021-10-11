package domain

type ParticipantRepository interface {
	StoreParticipant(participant *Participant) error

	GetParticipantById(participantId string) (*Participant, error)
	ListParticipants() ([]Participant, error)
}
