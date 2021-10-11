package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func UseLogrus(container *dig.Container)  {
	var server *echo.Echo
	err := container.Invoke(func(echoServer *echo.Echo) {
		server = echoServer
	})
	if err != nil {
		panic(err)
	}

	// Configure core
	logger := logrus.New()
	server.Logger = NewLogrusAdapter(logger)

	// Add logger to the dependency injection
	err = container.Provide(func() echo.Logger {
		return server.Logger
	})
	if err != nil {
		server.Logger.Fatal(err)
	}
}