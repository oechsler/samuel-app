package domain

type EventRepository interface {
	StoreEvent(event *Event) error

	GetEventById(eventId string) (*Event, error)
	ListEvents() ([]Event, error)
}
