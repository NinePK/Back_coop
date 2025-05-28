package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func MoobanRoutes(app *fiber.App) {
	Mooban_route := app.Group("/mooban")

	Mooban_route.Get("/", controllers.GetMoobans) // SELECT all moobans
	Mooban_route.Get("/amphur_id-:id", controllers.GetMoonbanByTambonId)
	Mooban_route.Get("/:id", controllers.GetMooban)     // SELECT mooban by ID
	Mooban_route.Post("/", controllers.CreateMooban)    // INSERT new mooban
	Mooban_route.Post("/:id", controllers.UpdateMooban) // UPDATE mooban by ID
	Mooban_route.Post("/:id", controllers.DeleteMooban) // DELETE mooban by ID
}
