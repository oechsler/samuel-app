package application

import (
	"github.com/cozy-hosting/messenger"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	core "github.com/oechsler/samuel-app-event/core/application"
	"github.com/oechsler/samuel-app-event/event/domain"
	"go.uber.org/dig"
)

type EventStartedHandler struct {
	dig.In

	Logger echo.Logger
	EventStore core.EventStore
	Repository domain.EventRepository
}

func NewEventStartedHandler(bus EventBus, logger echo.Logger, handler EventStartedHandler) {
	if err := bus.Consume(&domain.EventStarted{}, handler.handle); err != nil {
		logger.Fatal(err)
	}
}

func (h EventStartedHandler) handle(ctx messenger.Context) {
	delivery := ctx.GetDelivery()
	event, err := delivery.GetMessage()
	if err != nil {
		h.Logger.Error(err)
	}
	if err := h.EventStore.Store(event); err != nil {
		h.Logger.Error(err)
	}

	eventStarted := &domain.EventStarted{}
	if err := mapstructure.Decode(event.Body, eventStarted); err != nil {
		h.Logger.Error(err)
	}

	eventEntity, err := h.Repository.GetEventById(eventStarted.Id)
	if err != nil {
		h.Logger.Error(err)
	}

	eventEntity = eventEntity.Started(eventStarted.StartedAt)
	if err := h.Repository.StoreEvent(eventEntity); err != nil {
		h.Logger.Error(err)
	}

	if err := delivery.Acknowledge(); err != nil {
		h.Logger.Error(err)
	}
}