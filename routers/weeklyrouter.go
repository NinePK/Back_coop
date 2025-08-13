package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func WeeklyRoutes(app *fiber.App) {
	weekly_route := app.Group("/weekly")

	weekly_route.Get("/", controllers.GetWeeklys)                    // SELECT all Weekly reports
	weekly_route.Get("/:id", controllers.GetWeekly)                  // SELECT Weekly by ID
	weekly_route.Post("/", controllers.CreateWeekly)                 // INSERT new Weekly report
	weekly_route.Put("/:id", controllers.UpdateWeekly)               // UPDATE Weekly by ID
	weekly_route.Delete("/:id", controllers.DeleteWeekly)            // DELETE Weekly by ID
	weekly_route.Get("/training/:training_id", controllers.GetWeeklysByTraining) // GET Weekly reports by training ID
}