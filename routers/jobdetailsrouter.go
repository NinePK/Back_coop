package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func JobDetailsRoutes(app *fiber.App) {
	jobDetails := app.Group("/jobdetails")

	jobDetails.Post("/", controllers.CreateJobDetails)
	jobDetails.Get("/", controllers.GetAllJobDetails)
	jobDetails.Get("/training/:trainingId", controllers.GetJobDetailsByTrainingID)
	jobDetails.Put("/:id", controllers.UpdateJobDetails)
	jobDetails.Delete("/:id", controllers.DeleteJobDetails)
}