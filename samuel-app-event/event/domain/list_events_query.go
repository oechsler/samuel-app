package domain

type ListEvents struct {
	Skip int `query:"skip" validate:"gte=0"`
	Take int `query:"take" validate:"gt=0"`
}

func NewListEvents() *ListEvents {
	return &ListEvents{Skip: 0, Take: 10}
}