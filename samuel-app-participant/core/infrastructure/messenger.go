package infrastructure

import (
	"github.com/cozy-hosting/messenger"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"os"
)

func UseRabbitMessenger(container *dig.Container) func() {
	var logger echo.Logger
	if err := container.Invoke(func (echoLogger echo.Logger) {
		logger = echoLogger
	}); err != nil {
		panic(err)
	}

	connectionString := os.Getenv("RABBIT_MQ_CONNECTION")
	msgr, err := messenger.NewRabbitMessenger(connectionString)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Successfully connected to RabbitMQ")

	if err := container.Provide(func() messenger.Messenger {
		return msgr
	}); err != nil {
		logger.Fatal(err)
	}

	return func() {
		msgr.Close(func(err error) {
			logger.Fatal(err)
		})
	}
}
