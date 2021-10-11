package event

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/samuel-app-participant/event/infrastructure"
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
	}
	for _, providable := range providables {
		if err := container.Provide(providable); err != nil {
			logger.Fatal(err)
		}
	}

	logger.Info("Successfully loaded the Event module")
}
