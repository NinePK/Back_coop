package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func TeacherRoutes(app *fiber.App) {
	teacher_route := app.Group("/teacher")

	// Get students supervised by teacher
	teacher_route.Get("/students/:teacher_id", controllers.GetTeacherStudents)
	
	// Get teacher dashboard statistics
	teacher_route.Get("/:teacher_id/stats", controllers.GetTeacherStats)
	
	// Get all student reports for teacher
	teacher_route.Get("/:teacher_id/reports", controllers.GetTeacherStudentReports)
	
	// Get specific student's reports for teacher
	teacher_route.Get("/:teacher_id/students/:student_id/reports", controllers.GetStudentReportsForTeacher)
	
	// Update report status by teacher
	teacher_route.Put("/:teacher_id/reports/:report_id", controllers.UpdateReportStatusByTeacher)
}