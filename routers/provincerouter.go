package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProvinceRoutes(app *fiber.App) {
	Province_route := app.Group("/province")

	Province_route.Get("/", controllers.GetProvinces)      // SELECT all provinces
	Province_route.Get("/:id", controllers.GetProvince)       // SELECT province by ID
	Province_route.Post("/", controllers.CreateProvince)      // INSERT new province
	Province_route.Post("/:id", controllers.UpdateProvince)   // UPDATE province by ID
	Province_route.Post("/:id", controllers.DeleteProvince)   // DELETE province by ID
}