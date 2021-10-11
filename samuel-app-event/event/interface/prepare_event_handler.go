package _interface

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-event/event/application"
	"github.com/oechsler/samuel-app-event/event/domain"
	"go.uber.org/dig"
	"net/http"
)

type PrepareEventHandler struct {
	dig.In

	EventBus application.EventBus
}

func NewPrepareEventHandler(echo *echo.Echo, handler PrepareEventHandler) {
	group := echo.Group("/event")
	group.POST("", handler.handle)
}

func (h PrepareEventHandler) handle(ctx echo.Context) error {
	command := &domain.PrepareEvent{}
	if err := (&echo.DefaultBinder{}).BindBody(ctx, command); err != nil {
		return err
	}
	if err := validator.New().Struct(command); err != nil {
		return err
	}

	eventPrepared := domain.NewEventPrepared(command.Name)
	if err := h.EventBus.Publish(eventPrepared); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusAccepted)
}