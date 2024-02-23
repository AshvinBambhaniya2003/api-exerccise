package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/AshvinBambhaniya2003/titleandcredit/handlers"
)

func main() {
	// Create a new Fiber instance

	appConfig := fiber.Config{
		AppName:           "Online Streaming Plateform",
		EnablePrintRoutes: true,
		ServerHeader:      "Awesome App 1",
		CaseSensitive:     true,
	}

	app := fiber.New(appConfig)

	// Define routes
	app.Post("/titles", handlers.CreateTitle)
	app.Get("/titles", handlers.ListTitle)
	app.Put("/titles/:id", handlers.UpdateTitle)
	app.Delete("/titles/:id", handlers.DeleteTitle)

	// Start the server
	app.Listen(":3000")
}
