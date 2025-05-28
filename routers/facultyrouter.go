package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func FacultyRoutes(app *fiber.App) {
	Faculty_route := app.Group("/faculty")

	Faculty_route.Get("/", controllers.GetFacultys)      // SELECT all facultys
	Faculty_route.Get("/:id", controllers.GetFaculty)       // SELECT faculty by ID
	Faculty_route.Post("/", controllers.CreateFaculty)      // INSERT new faculty
	Faculty_route.Post("/:id", controllers.UpdateFaculty)   // UPDATE faculty by ID
	Faculty_route.Post("/:id", controllers.DeleteFaculty)   // DELETE faculty by ID
}