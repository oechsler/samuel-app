package main

import (
	"github.com/labstack/echo/v4"
	core "github.com/oechsler/samuel-app-event/core/infrastructure"
	"github.com/oechsler/samuel-app-event/event"
	"github.com/oechsler/samuel-app-event/participant"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	server := echo.New()

	// Make the server available for dependency injection
	if err := container.Provide(func() *echo.Echo {
		return server
	}); err != nil {
		panic(err)
	}

	// Register core modules
	core.UseLogrus(container)
	core.UseEnvironment(container)

	freeDatabase := core.UseMongoDatabase(container)
	defer freeDatabase()

	freeMessenger := core.UseRabbitMessenger(container)
	defer freeMessenger()

	core.UseEventStore(container)

	// Register first-party modules
	participant.UseModule(container)
	event.UseModule(container)

	// Start the server
	if err := server.Start(":1323"); err != nil {
		server.Logger.Fatal(err)
	}
}
