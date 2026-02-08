package routes

import (
	"PeepL-Test/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App){
	app.Post("/client", handlers.CreateClient)
	app.Put("/client/:id", handlers.UpdateClient)
	app.Delete("/client/:id", handlers.DeleteClient)
}