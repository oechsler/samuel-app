package event

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-event/event/application"
	"github.com/oechsler/samuel-app-event/event/infrastructure"
	_interface "github.com/oechsler/samuel-app-event/event/interface"
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
		infrastructure.NewEventRepositoryImpl,
		infrastructure.NewEventBusImpl,
	}
	for _, providable := range providables {
		if err := container.Provide(providable); err != nil {
			logger.Fatal(err)
		}
	}

	invokables := []interface{}{
		application.NewEventPreparedHandler,
		application.NewEventStartedHandler,
		application.NewEventCompletedHandler,

		_interface.NewPrepareEventHandler,
		_interface.NewStartEventHandler,
		_interface.NewCompleteEventHandler,

		_interface.NewListEventsHandler,
		_interface.NewRetrieveEventHandler,
	}
	for _, invokable := range invokables {
		if err := container.Invoke(invokable); err != nil {
			logger.Fatal(err)
		}
	}

	logger.Info("Successfully loaded the Event module")
}
