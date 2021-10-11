package domain

type ParticipantCheckedOut struct {
	Id string `param:"id"`
}

func NewParticipantCheckedOut(id string) *ParticipantCheckedOut {
	return &ParticipantCheckedOut{Id: id}
}
