package domain

type CheckOutParticipant struct {
	Id string `param:"id" validate:"required,uuid4"`
}
