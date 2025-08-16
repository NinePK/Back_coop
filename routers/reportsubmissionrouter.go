package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func ReportSubmissionRoutes(app *fiber.App) {
	reportSubmission := app.Group("/reportsubmission")

	reportSubmission.Post("/", controllers.CreateReportSubmission)
	reportSubmission.Get("/", controllers.GetAllReportSubmissions)
	reportSubmission.Get("/training/:trainingId", controllers.GetReportSubmissionByTrainingID)
	reportSubmission.Put("/:id", controllers.UpdateReportSubmission)
	reportSubmission.Put("/:id/approve", controllers.ApproveReportSubmission)
	reportSubmission.Delete("/:id", controllers.DeleteReportSubmission)
}