package controllers

import (
	"strconv"

	"coop_back/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReportSubmission(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var reportSubmission models.ReportSubmission
	
	if err := c.BodyParser(&reportSubmission); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// ตรวจสอบว่า training_id มีอยู่จริง
	var training models.Training
	if err := db.First(&training, reportSubmission.TrainingID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}

	// สร้าง record ใหม่
	if err := db.Create(&reportSubmission).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create report submission"})
	}

	// ดึงข้อมูลพร้อม relation
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&reportSubmission, reportSubmission.ID)
	
	// ส่งการแจ้งเตือนไปยังอาจารย์ที่ปรึกษา
	if training.TeacherID1 > 0 {
		if err := CreateDocumentNotification(db, training.UserID, training.TeacherID1, "coop10", reportSubmission.ReportTitleThai, reportSubmission.ID); err != nil {
			// Log error แต่ไม่ return error เพื่อไม่ให้กระทบการบันทึกเอกสาร
			// log.Printf("Failed to create notification: %v", err)
		}
	}
	
	return c.Status(201).JSON(reportSubmission)
}

func GetReportSubmissionByTrainingID(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	trainingID := c.Params("trainingId")
	
	var reportSubmission models.ReportSubmission
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Where("training_id = ?", trainingID).First(&reportSubmission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Report submission not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(reportSubmission)
}

func UpdateReportSubmission(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	reportSubmissionID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var reportSubmission models.ReportSubmission
	if err := db.First(&reportSubmission, reportSubmissionID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Report submission not found"})
	}

	var updateData models.ReportSubmission
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// อัปเดตข้อมูล
	if err := db.Model(&reportSubmission).Updates(&updateData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update report submission"})
	}

	// ดึงข้อมูลที่อัปเดตแล้ว
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&reportSubmission, reportSubmissionID)
	
	return c.JSON(reportSubmission)
}

func ApproveReportSubmission(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	reportSubmissionID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var reportSubmission models.ReportSubmission
	if err := db.First(&reportSubmission, reportSubmissionID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Report submission not found"})
	}

	// อัปเดตสถานะการอนุมัติ
	var approvalData struct {
		AdvisorApprovalStatus string `json:"advisorApprovalStatus"`
		AdvisorApprovalDate   string `json:"advisorApprovalDate"`
	}

	if err := c.BodyParser(&approvalData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	reportSubmission.AdvisorApprovalStatus = approvalData.AdvisorApprovalStatus
	reportSubmission.AdvisorApprovalDate = approvalData.AdvisorApprovalDate

	if err := db.Save(&reportSubmission).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update approval status"})
	}

	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&reportSubmission, reportSubmissionID)
	
	return c.JSON(reportSubmission)
}

func DeleteReportSubmission(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	reportSubmissionID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := db.Delete(&models.ReportSubmission{}, reportSubmissionID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete report submission"})
	}

	return c.JSON(fiber.Map{"message": "Report submission deleted successfully"})
}

func GetAllReportSubmissions(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var reportSubmissions []models.ReportSubmission
	
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Find(&reportSubmissions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(reportSubmissions)
}