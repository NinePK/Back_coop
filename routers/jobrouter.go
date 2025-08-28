package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func JobRoutes(app *fiber.App) {
	Job_route := app.Group("/job")

	Job_route.Get("/", controllers.GetJobs)            // SELECT all jobs
	Job_route.Get("/search", controllers.GetJobOption) // INSERT new user
	Job_route.Get("/:id", controllers.GetJob)          // SELECT job by ID
	Job_route.Post("/", controllers.CreateJob)         // INSERT new job
	Job_route.Post("/:id", controllers.UpdateJob)      // UPDATE job by ID
	Job_route.Post("/:id", controllers.DeleteJob)      // DELETE job by ID
}
