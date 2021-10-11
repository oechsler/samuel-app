package domain

type CompleteEvent struct {
	Id string `param:"id" validate:"required,uuid4"`
}