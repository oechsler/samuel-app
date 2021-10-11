package infrastructure

import (
	"github.com/cozy-hosting/clerk"
	"github.com/oechsler/samuel-app-participant/event/domain"
	"os"
)

type EventRepositoryImpl struct {
	connection clerk.Connection
	collection clerk.Collection
}

func NewEventRepositoryImpl(connection clerk.Connection) domain.EventRepository {
	databaseName := os.Getenv("MONGO_DB_DATABASE")
	collection := clerk.NewDatabase(databaseName).GetCollection("events")

	return &EventRepositoryImpl{connection: connection, collection: collection}
}

func (e EventRepositoryImpl) StoreEvent(event *domain.Event) error {
	storeCommand := clerk.NewMongoUpdateCommand(e.collection, event).
		Where("_id", event.Id).
		WithUpsert()

	return e.connection.SendCommand(storeCommand)
}

func (e EventRepositoryImpl) GetEventById(eventId string) (*domain.Event, error) {
	getByIdQuery := clerk.NewMongoSingleQuery(e.collection).Where("_id", eventId)

	iterator, err := e.connection.SendQuery(getByIdQuery)
	if err != nil {
		return nil, err
	}

	var event *domain.Event
	if err := iterator.Single(&event); err != nil {
		return nil, err
	}

	return event, nil
}

func (e EventRepositoryImpl) ListEvents() ([]domain.Event, error) {
	listQuery := clerk.NewMongoListQuery(e.collection)

	iterator, err := e.connection.SendQuery(listQuery)
	if err != nil {
		return nil, err
	}

	var events []domain.Event
	for iterator.Next() {
		var event domain.Event
		if err := iterator.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

