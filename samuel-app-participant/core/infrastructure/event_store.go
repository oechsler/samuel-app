package infrastructure

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func UseEventStore(container *dig.Container) {
	var logger echo.Logger
	if err := container.Invoke(func (echoLogger echo.Logger) {
		logger = echoLogger
	}); err != nil {
		panic(err)
	}

	if err := container.Provide(NewEventStoreImpl); err != nil {
		logger.Fatal(err)
	}

	logger.Info("Successfully provided the Event-Store")
}
