package infrastructure

import (
	"github.com/cozy-hosting/clerk"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"os"
)

func UseMongoDatabase(container *dig.Container) func() {
	var logger echo.Logger
	if err := container.Invoke(func (echoLogger echo.Logger) {
		logger = echoLogger
	}); err != nil {
		panic(err)
	}

	connectionString := os.Getenv("MONGO_DB_CONNECTION")
	connection, err := clerk.NewMongoConnection(connectionString)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Successfully connected to MongoDB")

	if err := container.Provide(func() clerk.Connection {
		return connection
	}); err != nil {
		logger.Fatal(err)
	}

	return func() {
		connection.Close(func(err error) {
			logger.Fatal(err)
		})
	}
}
