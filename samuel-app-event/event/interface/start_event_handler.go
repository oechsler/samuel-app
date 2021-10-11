package _interface

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-event/event/application"
	"github.com/oechsler/samuel-app-event/event/domain"
	"go.uber.org/dig"
	"net/http"
)

type StartEventHandler struct {
	dig.In

	EventBus application.EventBus
}

func NewStartEventHandler(echo *echo.Echo, handler StartEventHandler) {
	group := echo.Group("/event")
	group.POST("/:id/start", handler.handle)
}

func (h StartEventHandler) handle(ctx echo.Context) error {
	command := &domain.StartEvent{}
	_ = (&echo.DefaultBinder{}).BindPathParams(ctx, command)
	if err := validator.New().Struct(command); err != nil {
		return err
	}

	eventStarted := domain.NewEventStarted(command.Id)
	if err := h.EventBus.Publish(eventStarted); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusAccepted)
}