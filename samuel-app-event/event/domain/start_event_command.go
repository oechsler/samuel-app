package domain

type StartEvent struct {
	Id string `param:"id" validation:"required,uuid4"`
}
