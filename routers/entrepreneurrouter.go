package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func EntrepreneurRoutes(app *fiber.App) {
	Entrepreneur_route := app.Group("/entrepreneur")

	Entrepreneur_route.Get("/", controllers.GetEntrepreneurs)              // SELECT all entrepreneurs
	Entrepreneur_route.Get("/search", controllers.GetEntrepreneurOption)   // INSERT new user
	Entrepreneur_route.Get("/:id", controllers.GetEntrepreneur)            // SELECT entrepreneur by ID
	Entrepreneur_route.Post("/", controllers.CreateEntrepreneur)           // INSERT new entrepreneur
	Entrepreneur_route.Post("/:id", controllers.UpdateEntrepreneur)        // UPDATE entrepreneur by ID
	Entrepreneur_route.Post("/delete/:id", controllers.DeleteEntrepreneur) // DELETE entrepreneur by ID

}
