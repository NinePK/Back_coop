package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func RecordRoutes(app *fiber.App) {
	Record_route := app.Group("/record")

	Record_route.Get("/", controllers.GetRecords)          // SELECT all Records
	Record_route.Get("/:id", controllers.GetRecord)        // SELECT Record by ID
	Record_route.Post("/", controllers.CreateRecord)       // INSERT new Record
	Record_route.Post("/update", controllers.UpdateRecord) // UPDATE Record by ID
	Record_route.Post("/delete", controllers.DeleteRecord) // DELETE Record by ID

	// Record_route.Post("/search", controllers.SearchRecord) // INSERT new Record
}
