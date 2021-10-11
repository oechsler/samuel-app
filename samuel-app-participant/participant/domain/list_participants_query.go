package domain

type ListParticipants struct {
	Skip int `query:"skip" validate:"gte=0"`
	Take int `query:"take" validate:"gt=0"`
}

func NewListEvents() *ListParticipants {
	return &ListParticipants{Skip: 0, Take: 10}
}