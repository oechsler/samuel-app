package _interface

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-event/event/domain"
	"go.uber.org/dig"
	"net/http"
)

type ListEventsHandler struct {
	dig.In

	Repository domain.EventRepository
}

func NewListEventsHandler(echo *echo.Echo, handler ListEventsHandler) {
	group := echo.Group("/event")
	group.GET("", handler.handle)
}

func (h ListEventsHandler) handle(ctx echo.Context) error {
	query := domain.NewListEvents()
	_ = (&echo.DefaultBinder{}).BindQueryParams(ctx, query)
	if err := validator.New().Struct(query); err != nil {
		return err
	}

	events, err := h.Repository.ListEvents()
	if err != nil {
		return err
	}

	var filteredEvents []domain.Event
	linq.From(events).Skip(query.Skip).Take(query.Take).ToSlice(&filteredEvents)
	return ctx.JSON(http.StatusOK, filteredEvents)
}