package application

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/cozy-hosting/messenger"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/oechsler/samuel-app-participant/core/application"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"go.uber.org/dig"
)

type EveryParticipantCheckedOutHandler struct {
	dig.In

	Logger echo.Logger
	EventStore application.EventStore
	Repository domain.ParticipantRepository
}

func NewParticipantCheckOutEveryoneHandler (bus ParticipantBus, logger echo.Logger, handler EveryParticipantCheckedOutHandler) {
	if err := bus.Consume(&domain.EveryParticipantCheckedOut{}, handler.handle); err != nil {
		logger.Fatal(err)
	}
}

func (h EveryParticipantCheckedOutHandler) handle(ctx messenger.Context) {
	delivery := ctx.GetDelivery()
	event, err := delivery.GetMessage()
	if err != nil {
		h.Logger.Error(err)
	}
	if err := h.EventStore.Store(event); err != nil {
		h.Logger.Error(err)
	}

	everyParticipantCheckedOut := &domain.EveryParticipantCheckedOut{}
	if err := mapstructure.Decode(event.Body, everyParticipantCheckedOut); err != nil {
		h.Logger.Error(err)
	}

	participantEntities, err := h.Repository.ListParticipants()
	if err != nil {
		h.Logger.Error(err)
	}

	linq.From(participantEntities).WhereT(func(participant domain.Participant) bool {
		return participant.EventId == everyParticipantCheckedOut.EventId
	}).ForEachT(func(participant domain.Participant) {
		checkedOutParticipant := participant.CheckOut()

		if err := h.Repository.StoreParticipant(checkedOutParticipant); err != nil {
			h.Logger.Error(err)
		}
	})

	if err := delivery.Acknowledge(); err != nil {
		h.Logger.Error(err)
	}
}