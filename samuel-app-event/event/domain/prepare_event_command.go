package domain

type PrepareEvent struct {
	Name string `json:"name" validate:"required"`
}
