package infrastructure

import (
	"github.com/cozy-hosting/messenger"
	core "github.com/oechsler/samuel-app-event/core/application"
	"github.com/oechsler/samuel-app-event/participant/application"
	"os"
	"reflect"
)

type ParticipantBusImpl struct {
	exchange messenger.Exchange
	msgr messenger.Messenger

	eventStore core.EventStore
}

func NewParticipantBusImpl(msgr messenger.Messenger, eventStore core.EventStore) application.ParticipantBus {
	exchangeName := os.Getenv("RABBIT_MQ_EXCHANGE")
	exchange := messenger.NewExchange().Named(exchangeName)

	return &ParticipantBusImpl{
		exchange: exchange,
		msgr: msgr,
		eventStore: eventStore,
	}
}

func (p ParticipantBusImpl) Publish(event interface{}) error {
	eventTypeString := reflect.TypeOf(event).String()
	queueToPublishOn := messenger.NewQueue().
		WithTopic(eventTypeString)

	messageToPublish := messenger.NewMessage(event)

	if err := p.eventStore.Store(messageToPublish); err != nil {
		return err
	}

	return p.msgr.Publish(p.exchange, queueToPublishOn, messageToPublish)
}

func (p ParticipantBusImpl) Consume(eventType interface{}, handler func(ctx messenger.Context)) error {
	eventTypeString := reflect.TypeOf(eventType).String()
	queueToConsumeOn := messenger.NewQueue().
		Named(eventTypeString)

	consumer := messenger.NewConsumer(handler)
	if free, err := p.msgr.Consume(p.exchange, queueToConsumeOn, consumer); err != nil {
		_ = free()
		return err
	}

	return nil
}
