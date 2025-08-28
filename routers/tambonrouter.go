package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func TambonRoutes(app *fiber.App) {
	Tambon_route := app.Group("/tambon")

	Tambon_route.Get("/", controllers.GetTambons) // SELECT all tambons
	Tambon_route.Get("/amphur_id-:id", controllers.GetTambonByAmphurId)
	Tambon_route.Get("/:id", controllers.GetTambon)     // SELECT tambon by ID
	Tambon_route.Post("/", controllers.CreateTambon)    // INSERT new tambon
	Tambon_route.Post("/:id", controllers.UpdateTambon) // UPDATE tambon by ID
	Tambon_route.Post("/:id", controllers.DeleteTambon) // DELETE tambon by ID
}
