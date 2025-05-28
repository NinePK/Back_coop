package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func InchargeRoutes(app *fiber.App) {
	Incharge_route := app.Group("/incharge")

	Incharge_route.Get("/", controllers.GetIncharges)      // SELECT all incharges
	Incharge_route.Get("/:id", controllers.GetIncharge)       // SELECT incharge by ID
	Incharge_route.Post("/", controllers.CreateIncharge)      // INSERT new incharge
	Incharge_route.Post("/:id", controllers.UpdateIncharge)   // UPDATE incharge by ID
	Incharge_route.Post("/:id", controllers.DeleteIncharge)   // DELETE incharge by ID
}