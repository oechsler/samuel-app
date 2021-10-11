package application

import (
	"github.com/cozy-hosting/messenger"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/oechsler/samuel-app-participant/core/application"
	eventDomain "github.com/oechsler/samuel-app-participant/event/domain"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"go.uber.org/dig"
)

type ParticipantCheckedInHandler struct {
	dig.In

	Logger echo.Logger
	EventStore application.EventStore
	EventRepository eventDomain.EventRepository
	ParticipantRepository domain.ParticipantRepository
}

func NewParticipantCheckedInHandler(bus ParticipantBus, logger echo.Logger, handler ParticipantCheckedInHandler) {
	if err := bus.Consume(&domain.ParticipantCheckedIn{}, handler.handle); err != nil {
		logger.Fatal(err)
	}
}

func (h ParticipantCheckedInHandler) handle(ctx messenger.Context) {
	delivery := ctx.GetDelivery()
	event, err := delivery.GetMessage()
	if err != nil {
		h.Logger.Error(err)
	}
	if err := h.EventStore.Store(event); err != nil {
		h.Logger.Error(err)
	}

	participantCheckedIn := &domain.ParticipantCheckedIn{}
	if err := mapstructure.Decode(event.Body, participantCheckedIn); err != nil {
		h.Logger.Error(err)
	}

	eventEntity, err := h.EventRepository.GetEventById(participantCheckedIn.EventId)
	if err != nil {
		h.Logger.Infof("Event %s does not exist", participantCheckedIn.EventId)
		if err := delivery.Acknowledge(); err != nil {
			h.Logger.Error(err)
		}
		return
	}
	if !eventEntity.IsActive() {
		h.Logger.Infof("Event %s is not active", participantCheckedIn.EventId)
		if err := delivery.Acknowledge(); err != nil {
			h.Logger.Error(err)
		}
		return
	}

	participantEntity := domain.NewParticipant(
		participantCheckedIn.Id,
		participantCheckedIn.EventId,
		participantCheckedIn.Address,
	)
	if err := h.ParticipantRepository.StoreParticipant(participantEntity); err != nil {
		h.Logger.Error(err)
	}

	if err := delivery.Acknowledge(); err != nil {
		h.Logger.Error(err)
	}
}
