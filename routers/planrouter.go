package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func PlanRoutes(app *fiber.App) {
	plan_route := app.Group("/plan")

	plan_route.Get("/", controllers.GetPlans)      // SELECT all plans
	plan_route.Get("/:id", controllers.GetPlan)       // SELECT plan by ID
	plan_route.Post("/", controllers.CreatePlan)      // INSERT new plan
	plan_route.Post("/:id", controllers.UpdatePlan)   // UPDATE plan by ID
	plan_route.Post("/:id", controllers.DeletePlan)   // DELETE plan by ID
}