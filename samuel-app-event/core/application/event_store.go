package application

import "github.com/cozy-hosting/messenger"

type EventStore interface {
	Store(event messenger.Message) error
}
