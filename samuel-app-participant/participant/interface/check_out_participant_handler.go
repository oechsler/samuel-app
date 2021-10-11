package _interface

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-participant/participant/application"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"go.uber.org/dig"
	"net/http"
)

type CheckOutParticipantHandler struct {
	dig.In

	ParticipantBus application.ParticipantBus
}

func NewCheckOutParticipantHandler(echo *echo.Echo, handler CheckOutParticipantHandler) {
	group := echo.Group("/participant")
	group.POST("/:id/checkout", handler.handle)
}

func (h CheckOutParticipantHandler) handle(ctx echo.Context) error {
	command := &domain.CheckOutParticipant{}
	_ = (&echo.DefaultBinder{}).BindPathParams(ctx, command)
	if err := validator.New().Struct(command); err != nil {
		return err
	}

	participantCheckedOut := domain.NewParticipantCheckedOut(command.Id)
	if err := h.ParticipantBus.Publish(participantCheckedOut); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusAccepted)
}
