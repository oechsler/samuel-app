package infrastructure

import (
	"github.com/cozy-hosting/messenger"
	core "github.com/oechsler/samuel-app-event/core/application"
	"github.com/oechsler/samuel-app-event/event/application"
	"os"
	"reflect"
)

type EventBusImpl struct {
	exchange messenger.Exchange
	msgr messenger.Messenger

	eventStore core.EventStore
}

func NewEventBusImpl(msgr messenger.Messenger, eventStore core.EventStore) application.EventBus {
	exchangeName := os.Getenv("RABBIT_MQ_EXCHANGE")
	exchange := messenger.NewExchange().Named(exchangeName)

	return &EventBusImpl{
		exchange: exchange,
		msgr: msgr,
		eventStore: eventStore,
	}
}

func (e EventBusImpl) Publish(event interface{}) error {
	eventTypeString := reflect.TypeOf(event).String()
	queueToPublishOn := messenger.NewQueue().
		WithTopic(eventTypeString)

	messageToPublish := messenger.NewMessage(event)

	if err := e.eventStore.Store(messageToPublish); err != nil {
		return err
	}

	return e.msgr.Publish(e.exchange, queueToPublishOn, messageToPublish)
}

func (e EventBusImpl) Consume(eventType interface{}, handler func(ctx messenger.Context)) error {
	eventTypeString := reflect.TypeOf(eventType).String()
	queueToConsumeOn := messenger.NewQueue().
		Named(eventTypeString)

	consumer := messenger.NewConsumer(handler)
	if free, err := e.msgr.Consume(e.exchange, queueToConsumeOn, consumer); err != nil {
		_ = free()
		return err
	}

	return nil
}

