package controllers

import (
	"coop_back/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get students supervised by a teacher
func GetTeacherStudents(c *fiber.Ctx) error {
	teacherID := c.Params("teacher_id")
	db := c.Locals("db").(*gorm.DB)

	var trainings []models.Training

	// ดึงข้อมูลการฝึกงานของนิสิตที่อาจารย์ดูแล
	if result := db.
		Preload("User").
		Preload("User.Major").
		Preload("User.Major.Faculty").
		Preload("Job").
		Preload("Job.Entrepreneur").
		Where("teacher_id1 = ? OR teacher_id2 = ?", teacherID, teacherID).
		Find(&trainings); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch teacher's students",
		})
	}

	return c.JSON(trainings)
}

// Get teacher's dashboard statistics
func GetTeacherStats(c *fiber.Ctx) error {
	teacherID := c.Params("teacher_id")
	db := c.Locals("db").(*gorm.DB)

	// Convert teacherID to int64
	teacherIDInt, err := strconv.ParseInt(teacherID, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid teacher ID",
		})
	}

	stats := make(map[string]interface{})

	// นับจำนวนนิสิตทั้งหมดที่ดูแล
	var totalStudents int64
	db.Model(&models.Training{}).
		Where("teacher_id1 = ? OR teacher_id2 = ?", teacherIDInt, teacherIDInt).
		Count(&totalStudents)

	// นับจำนวนนิสิตที่กำลังฝึกงาน (สถานะ active)
	var activeInternships int64
	db.Model(&models.Training{}).
		Where("(teacher_id1 = ? OR teacher_id2 = ?) AND status = ?", teacherIDInt, teacherIDInt, 1).
		Count(&activeInternships)

	// นับจำนวนรายงานที่เสร็จสิ้น
	var completedReports int64
	db.Table("weekly").
		Joins("JOIN training ON weekly.training_id = training.id").
		Where("training.teacher_id1 = ? OR training.teacher_id2 = ?", teacherIDInt, teacherIDInt).
		Where("weekly.status = ?", "completed").
		Count(&completedReports)

	// นับจำนวนรายงานที่รออนุมัติ
	var pendingReports int64
	db.Table("weekly").
		Joins("JOIN training ON weekly.training_id = training.id").
		Where("training.teacher_id1 = ? OR training.teacher_id2 = ?", teacherIDInt, teacherIDInt).
		Where("weekly.status = ?", "pending").
		Count(&pendingReports)

	stats["totalStudents"] = totalStudents
	stats["activeInternships"] = activeInternships
	stats["completedReports"] = completedReports
	stats["pendingReports"] = pendingReports

	return c.JSON(stats)
}

// Get weekly reports of students supervised by teacher
func GetTeacherStudentReports(c *fiber.Ctx) error {
	teacherID := c.Params("teacher_id")
	db := c.Locals("db").(*gorm.DB)

	var reports []models.Weekly

	// ดึงรายงานประจำสัปดาห์ของนิสิตที่อาจารย์ดูแล
	if result := db.
		Preload("Training").
		Preload("Training.User").
		Preload("Training.Job").
		Preload("Training.Job.Entrepreneur").
		Joins("JOIN training ON weekly.training_id = training.id").
		Where("training.teacher_id1 = ? OR training.teacher_id2 = ?", teacherID, teacherID).
		Order("weekly.created_at DESC").
		Find(&reports); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch student reports",
		})
	}

	return c.JSON(reports)
}

// Get specific student's reports for teacher
func GetStudentReportsForTeacher(c *fiber.Ctx) error {
	teacherID := c.Params("teacher_id")
	studentID := c.Params("student_id")
	db := c.Locals("db").(*gorm.DB)

	var reports []models.Weekly

	// ตรวจสอบว่าอาจารย์คนนี้ดูแลนิสิตคนนี้หรือไม่
	var training models.Training
	if result := db.Where("user_id = ? AND (teacher_id1 = ? OR teacher_id2 = ?)", 
		studentID, teacherID, teacherID).First(&training); result.Error != nil {
		return c.Status(403).JSON(fiber.Map{
			"error": "Unauthorized: You don't supervise this student",
		})
	}

	// ดึงรายงานของนิสิตคนนี้
	if result := db.
		Preload("Training").
		Preload("Training.User").
		Preload("Training.Job").
		Preload("Training.Job.Entrepreneur").
		Where("training_id = ?", training.ID).
		Order("week ASC").
		Find(&reports); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch student reports",
		})
	}

	return c.JSON(reports)
}

// Update weekly report status by teacher
func UpdateReportStatusByTeacher(c *fiber.Ctx) error {
	teacherID := c.Params("teacher_id")
	reportID := c.Params("report_id")
	db := c.Locals("db").(*gorm.DB)

	// Parse request body
	var request struct {
		Status   string `json:"status"`
		Comment  string `json:"comment"`
		Grade    int    `json:"grade"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// ตรวจสอบว่ารายงานนี้เป็นของนิสิตที่อาจารย์ดูแลหรือไม่
	var report models.Weekly
	if result := db.
		Preload("Training").
		Where("id = ?", reportID).
		First(&report); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Report not found",
		})
	}

	// ตรวจสอบสิทธิ์
	if report.Training.TeacherID1 != parseInt64(teacherID) && 
	   report.Training.TeacherID2 != parseInt64(teacherID) {
		return c.Status(403).JSON(fiber.Map{
			"error": "Unauthorized: You don't supervise this student",
		})
	}

	// อัปเดตสถานะรายงาน
	report.Status = request.Status
	// Note: เพิ่มฟิลด์ comment และ grade ในโมเดล Weekly หากจำเป็น

	if result := db.Save(&report); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update report status",
		})
	}

	return c.JSON(report)
}

// Helper function to parse string to int64
func parseInt64(s string) int64 {
	if val, err := strconv.ParseInt(s, 10, 64); err == nil {
		return val
	}
	return 0
}