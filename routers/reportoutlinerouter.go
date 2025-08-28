package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func ReportOutlineRoutes(app *fiber.App) {
	reportOutline := app.Group("/reportoutline")

	reportOutline.Post("/", controllers.CreateReportOutline)
	reportOutline.Get("/", controllers.GetAllReportOutlines)
	reportOutline.Get("/training/:trainingId", controllers.GetReportOutlineByTrainingID)
	reportOutline.Put("/:id", controllers.UpdateReportOutline)
	reportOutline.Delete("/:id", controllers.DeleteReportOutline)
}