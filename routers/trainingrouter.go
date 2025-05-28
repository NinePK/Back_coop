package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func TrainingRoutes(app *fiber.App) {
	Training_route := app.Group("/training")

	Training_route.Get("/", controllers.GetTrainings)          // SELECT all Trainings
	Training_route.Get("/:id", controllers.GetTraining)        // SELECT Training by ID
	Training_route.Post("/", controllers.CreateTraining)       // INSERT new Training
	Training_route.Post("/update", controllers.UpdateTraining) // UPDATE Training by ID
	Training_route.Post("/delete", controllers.DeleteTraining) // DELETE Training by ID

	Training_route.Get("/user/:user_id<int>-:semester_id<int>", controllers.GetTrainingsByUser) // INSERT new Training
}
