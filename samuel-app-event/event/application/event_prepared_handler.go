package application

import (
	"github.com/cozy-hosting/messenger"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	core "github.com/oechsler/samuel-app-event/core/application"
	"github.com/oechsler/samuel-app-event/event/domain"
	"go.uber.org/dig"
)

type EventPreparedHandler struct {
	dig.In

	Logger echo.Logger
	EventStore core.EventStore
	Repository domain.EventRepository
}

func NewEventPreparedHandler(bus EventBus, logger echo.Logger, handler EventPreparedHandler) {
	if err := bus.Consume(&domain.EventPrepared{}, handler.handle); err != nil {
		logger.Fatal(err)
	}
}

func (h EventPreparedHandler) handle(ctx messenger.Context) {
	delivery := ctx.GetDelivery()
	event, err := delivery.GetMessage()
	if err != nil {
		h.Logger.Error(err)
	}
	if err := h.EventStore.Store(event); err != nil {
		h.Logger.Error(err)
	}

	var eventPrepared domain.EventPrepared
	if err := mapstructure.Decode(event.Body, &eventPrepared); err != nil {
		h.Logger.Error(err)
	}

	eventEntity := domain.NewEvent(eventPrepared.Id, eventPrepared.Name)
	if err := h.Repository.StoreEvent(eventEntity); err != nil {
		h.Logger.Error(err)
	}

	if err := delivery.Acknowledge(); err != nil {
		h.Logger.Error(err)
	}
}

