package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func MajorRoutes(app *fiber.App) {
	Major_route := app.Group("/major")

	Major_route.Get("/", controllers.GetMajors)   // SELECT all Majors
	Major_route.Get("/:id", controllers.GetMajor) // SELECT Major by ID
	Major_route.Get("/:param/:val", controllers.GetMajorWhere)
	Major_route.Post("/", controllers.CreateMajor)    // INSERT new Major
	Major_route.Post("/:id", controllers.UpdateMajor) // UPDATE Major by ID
	Major_route.Post("/:id", controllers.DeleteMajor) // DELETE Major by ID
}
