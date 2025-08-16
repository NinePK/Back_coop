package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func Coop04AccommodationRoutes(app *fiber.App) {
	// COOP-04 Accommodation routes
	app.Post("/coop04-accommodation/", controllers.CreateCoop04Accommodation)
	app.Get("/coop04-accommodation/", controllers.GetCoop04Accommodation)
	app.Put("/coop04-accommodation/:id", controllers.UpdateCoop04Accommodation)
	app.Post("/coop04-accommodation/update", controllers.UpdateCoop04Accommodation)
}