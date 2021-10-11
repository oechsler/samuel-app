package _interface

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"go.uber.org/dig"
	"net/http"
)

type RetrieveEventHandler struct {
	dig.In

	Repository domain.ParticipantRepository
}

func NewRetrieveParticipantHandler(echo *echo.Echo, handler RetrieveEventHandler) {
	group := echo.Group("/participant")
	group.GET("/:id", handler.handle)
}

func (h RetrieveEventHandler) handle(ctx echo.Context) error {
	query := &domain.RetrieveParticipant{}
	_ = (&echo.DefaultBinder{}).BindPathParams(ctx, query)
	if err := validator.New().Struct(query); err != nil {
		return err
	}

	participant, err := h.Repository.GetParticipantById(query.Id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, participant)
}