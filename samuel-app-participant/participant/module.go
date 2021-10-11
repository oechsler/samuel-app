package participant

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-participant/participant/application"
	"github.com/oechsler/samuel-app-participant/participant/infrastructure"
	_interface "github.com/oechsler/samuel-app-participant/participant/interface"
	"go.uber.org/dig"
)

func UseModule(container *dig.Container) {
	var logger echo.Logger
	if err := container.Invoke(func (echoLogger echo.Logger) {
		logger = echoLogger
	}); err != nil {
		panic(err)
	}

	providables := []interface{}{
		infrastructure.NewParticipantRepositoryImpl,
		infrastructure.NewParticipantBusImpl,
	}
	for _, providable := range providables {
		if err := container.Provide(providable); err != nil {
			logger.Fatal(err)
		}
	}

	invokables := []interface{}{
		application.NewParticipantCheckedInHandler,
		application.NewParticipantCheckedOutHandler,
		application.NewParticipantCheckOutEveryoneHandler,

		_interface.NewCheckInParticipantHandler,
		_interface.NewCheckOutParticipantHandler,

		_interface.NewListParticipantsHandler,
		_interface.NewRetrieveParticipantHandler,
	}
	for _, invokable := range invokables {
		if err := container.Invoke(invokable); err != nil {
			logger.Fatal(err)
		}
	}

	logger.Info("Successfully loaded the Participant module")
}
