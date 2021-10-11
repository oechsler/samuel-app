package domain

type RetrieveEvent struct {
	Id string `param:"id" validate:"required"`
}
