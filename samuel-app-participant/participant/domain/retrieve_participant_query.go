package domain

type RetrieveParticipant struct {
	Id string `param:"id" validate:"required"`
}
