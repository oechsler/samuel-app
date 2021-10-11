package infrastructure

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"os"
)

func UseEnvironment(container *dig.Container) {
	var logger echo.Logger
	if err := container.Invoke(func (echoLogger echo.Logger) {
		logger = echoLogger
	}); err != nil {
		panic(err)
	}

	production := false
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		production = true
	}

	if !production {
		if err := godotenv.Load(); err != nil {
			logger.Warn("Development environment file could not be located")
		}

		logger.Info("Running in development mode")
		return
	}

	logger.Info("Running in production mode")
}
