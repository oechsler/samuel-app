package _interface

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"go.uber.org/dig"
	"net/http"
)

type ListParticipantsHandler struct {
	dig.In

	Repository domain.ParticipantRepository
}

func NewListParticipantsHandler(echo *echo.Echo, handler ListParticipantsHandler) {
	group := echo.Group("/participant")
	group.GET("", handler.handle)
}

func (h ListParticipantsHandler) handle(ctx echo.Context) error {
	query := domain.NewListEvents()
	_ = (&echo.DefaultBinder{}).BindQueryParams(ctx, query)
	if err := validator.New().Struct(query); err != nil {
		return err
	}

	participants, err := h.Repository.ListParticipants()
	if err != nil {
		return err
	}

	var filteredParticipants []domain.Participant
	linq.From(participants).Skip(query.Skip).Take(query.Take).ToSlice(&filteredParticipants)
	return ctx.JSON(http.StatusOK, filteredParticipants)
}