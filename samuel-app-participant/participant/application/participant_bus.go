package application

import "github.com/cozy-hosting/messenger"

type ParticipantBus interface {
	Publish(event interface{}) error
	Consume(eventType interface{}, handler func(ctx messenger.Context)) error
}
