package _interface

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-participant/participant/application"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"go.uber.org/dig"
	"net/http"
)

type CheckInParticipantHandler struct {
	dig.In

	ParticipantBus application.ParticipantBus
}

func NewCheckInParticipantHandler(echo *echo.Echo, handler CheckInParticipantHandler) {
	group := echo.Group("/participant")
	group.POST("/checkin", handler.handle)
}

func (h CheckInParticipantHandler) handle(ctx echo.Context) error {
	command := &domain.CheckInParticipant{}
	if err := (&echo.DefaultBinder{}).BindBody(ctx, command); err != nil {
		return err
	}
	if err := validator.New().Struct(command); err != nil {
		return err
	}

	participantAddress := domain.NewAddress(
		command.Address.Street,
		command.Address.Number,
		command.Address.ZipCode,
		command.Address.City,
	)
	participantCheckedIn := domain.NewParticipantCheckedIn(command.EventId, participantAddress)
	if err := h.ParticipantBus.Publish(participantCheckedIn); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusAccepted)
}
