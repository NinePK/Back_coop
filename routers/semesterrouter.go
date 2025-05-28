package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func SemesterRoutes(app *fiber.App) {
	Semester_route := app.Group("/semester")

	Semester_route.Get("/", controllers.GetSemesters)                // SELECT all semesters
	Semester_route.Get("/current", controllers.GetCurrentSemester)   // SELECT semester by ID
	Semester_route.Get("/:id", controllers.GetSemester)              // SELECT semester by ID
	Semester_route.Post("/", controllers.CreateSemester)             // INSERT new semester
	Semester_route.Post("/update", controllers.UpdateSemesterByData) // UPDATE semester by ID
	Semester_route.Post("/update/:id", controllers.UpdateSemester)   // UPDATE semester by ID
	Semester_route.Post("/delete/:id", controllers.DeleteSemester)   // DELETE semester by ID
}
