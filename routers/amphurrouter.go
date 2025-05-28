package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func AmphurRoutes(app *fiber.App) {
	Amphur_route := app.Group("/amphur")

	Amphur_route.Get("/", controllers.GetAmphurs)                           // SELECT all amphurs
	Amphur_route.Get("/province_id-:id", controllers.GetAmphurByProvinceId) // SELECT amphur by ID
	Amphur_route.Get("/:id", controllers.GetAmphur)                         // SELECT amphur by ID
	Amphur_route.Post("/", controllers.CreateAmphur)                        // INSERT new amphur
	Amphur_route.Post("/:id", controllers.UpdateAmphur)                     // UPDATE amphur by ID
	Amphur_route.Post("/:id", controllers.DeleteAmphur)                     // DELETE amphur by ID
}
