package _interface

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-event/event/application"
	"github.com/oechsler/samuel-app-event/event/domain"
	participantApplication "github.com/oechsler/samuel-app-event/participant/application"
	participantDomain "github.com/oechsler/samuel-app-event/participant/domain"
	"go.uber.org/dig"
	"net/http"
)

type CompleteEventHandler struct {
	dig.In

	EventBus application.EventBus
	ParticipantBus participantApplication.ParticipantBus
}

func NewCompleteEventHandler(echo *echo.Echo, handler CompleteEventHandler) {
	group := echo.Group("/event")
	group.POST("/:id/complete", handler.handle)
}

func (h CompleteEventHandler) handle(ctx echo.Context) error {
	command := &domain.CompleteEvent{}
	_ = (&echo.DefaultBinder{}).BindPathParams(ctx, command)
	if err := validator.New().Struct(command); err != nil {
		return err
	}

	everyParticipantCheckedOut := participantDomain.NewEveryParticipantCheckedOut(command.Id)
	if err := h.ParticipantBus.Publish(everyParticipantCheckedOut); err != nil {
		return err
	}

	eventCompleted := domain.NewEventCompleted(command.Id)
	if err := h.EventBus.Publish(eventCompleted); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusAccepted)
}
