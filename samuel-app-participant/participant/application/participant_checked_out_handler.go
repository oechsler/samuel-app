package application

import (
	"github.com/cozy-hosting/messenger"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/oechsler/samuel-app-participant/core/application"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"go.uber.org/dig"
)

type ParticipantCheckedOutHandler struct {
	dig.In

	Logger echo.Logger
	EventStore application.EventStore
	Repository domain.ParticipantRepository
}

func NewParticipantCheckedOutHandler(bus ParticipantBus, logger echo.Logger, handler ParticipantCheckedOutHandler) {
	if err := bus.Consume(&domain.ParticipantCheckedOut{}, handler.handle); err != nil {
		logger.Fatal(err)
	}
}

func (h ParticipantCheckedOutHandler) handle(ctx messenger.Context) {
	delivery := ctx.GetDelivery()
	event, err := delivery.GetMessage()
	if err != nil {
		h.Logger.Error(err)
	}
	if err := h.EventStore.Store(event); err != nil {
		h.Logger.Error(err)
	}

	participantCheckedOut := &domain.ParticipantCheckedOut{}
	if err := mapstructure.Decode(event.Body, participantCheckedOut); err != nil {
		h.Logger.Error(err)
	}

	participantEntity, err := h.Repository.GetParticipantById(participantCheckedOut.Id)
	if err != nil {
		h.Logger.Infof("Participant %s does not exists", participantCheckedOut.Id)
		if err := delivery.Acknowledge(); err != nil {
			h.Logger.Error(err)
		}
		return
	}

	participantEntity = participantEntity.CheckOut()
	if err := h.Repository.StoreParticipant(participantEntity); err != nil {
		h.Logger.Error(err)
	}

	if err := delivery.Acknowledge(); err != nil {
		h.Logger.Error(err)
	}
}