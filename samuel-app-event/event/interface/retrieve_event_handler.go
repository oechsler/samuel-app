package _interface

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-event/event/domain"
	"go.uber.org/dig"
	"net/http"
)

type RetrieveEventHandler struct {
	dig.In

	Repository domain.EventRepository
}

func NewRetrieveEventHandler(echo *echo.Echo, handler RetrieveEventHandler) {
	group := echo.Group("/event")
	group.GET("/:id", handler.handle)
}

func (h RetrieveEventHandler) handle(ctx echo.Context) error {
	query := &domain.RetrieveEvent{}
	_ = (&echo.DefaultBinder{}).BindPathParams(ctx, query)
	if err := validator.New().Struct(query); err != nil {
		return err
	}

	event, err := h.Repository.GetEventById(query.Id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, event)
}