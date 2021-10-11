package infrastructure

import (
	"github.com/cozy-hosting/clerk"
	"github.com/cozy-hosting/messenger"
	"github.com/oechsler/samuel-app-participant/core/application"
	"os"
)

type EventStoreImpl struct {
	connection clerk.Connection
	collection clerk.Collection
}

func NewEventStoreImpl(connection clerk.Connection) application.EventStore {
	databaseName := os.Getenv("MONGO_DB_DATABASE")
	collection := clerk.NewDatabase(databaseName).GetCollection("event-store")

	return &EventStoreImpl{connection: connection, collection: collection}
}

func (e EventStoreImpl) Store(event messenger.Message) error {
	storeCommand := clerk.NewMongoUpdateCommand(e.collection, event).
		Where("series", event.Series).
		Where("revision", event.Revision).
		WithUpsert()

	return e.connection.SendCommand(storeCommand)
}

