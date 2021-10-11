package application

import (
	"github.com/cozy-hosting/messenger"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	core "github.com/oechsler/samuel-app-event/core/application"
	"github.com/oechsler/samuel-app-event/event/domain"
	"go.uber.org/dig"
)

type EventCompletedHandler struct {
	dig.In

	Logger echo.Logger
	EventStore core.EventStore
	Repository domain.EventRepository
}

func NewEventCompletedHandler(bus EventBus, logger echo.Logger, handler EventCompletedHandler) {
	if err := bus.Consume(&domain.EventCompleted{}, handler.handle); err != nil {
		logger.Fatal(err)
	}
}

func (h EventCompletedHandler) handle(ctx messenger.Context) {
	delivery := ctx.GetDelivery()
	event, err := delivery.GetMessage()
	if err != nil {
		h.Logger.Error(err)
	}
	if err := h.EventStore.Store(event); err != nil {
		h.Logger.Error(err)
	}

	eventCompleted := &domain.EventCompleted{}
	if err := mapstructure.Decode(event.Body, eventCompleted); err != nil {
		h.Logger.Error(err)
	}

	eventEntity, err := h.Repository.GetEventById(eventCompleted.Id)
	if err != nil {
		h.Logger.Error(err)
	}

	eventEntity = eventEntity.Completed(eventCompleted.CompletedAt)
	if err := h.Repository.StoreEvent(eventEntity); err != nil {
		h.Logger.Error(err)
	}

	if err := delivery.Acknowledge(); err != nil {
		h.Logger.Error(err)
	}
}
