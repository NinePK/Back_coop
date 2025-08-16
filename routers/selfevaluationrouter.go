package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func SelfEvaluationRoutes(app *fiber.App) {
	selfEvaluation := app.Group("/selfevaluation")

	selfEvaluation.Post("/", controllers.CreateSelfEvaluation)
	selfEvaluation.Get("/", controllers.GetAllSelfEvaluations)
	selfEvaluation.Get("/training/:trainingId", controllers.GetSelfEvaluationByTrainingID)
	selfEvaluation.Get("/statistics", controllers.GetSelfEvaluationStatistics)
	selfEvaluation.Put("/:id", controllers.UpdateSelfEvaluation)
	selfEvaluation.Delete("/:id", controllers.DeleteSelfEvaluation)
}