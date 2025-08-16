package controllers

import (
	"strconv"

	"coop_back/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReportOutline(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var reportOutline models.ReportOutline
	
	if err := c.BodyParser(&reportOutline); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// ตรวจสอบว่า training_id มีอยู่จริง
	var training models.Training
	if err := db.First(&training, reportOutline.TrainingID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}

	// สร้าง record ใหม่
	if err := db.Create(&reportOutline).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create report outline"})
	}

	// ดึงข้อมูลพร้อม relation
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&reportOutline, reportOutline.ID)
	
	// ส่งการแจ้งเตือนไปยังอาจารย์ที่ปรึกษา
	if training.TeacherID1 > 0 {
		if err := CreateDocumentNotification(db, training.UserID, training.TeacherID1, "coop07", "", reportOutline.ID); err != nil {
			// Log error แต่ไม่ return error เพื่อไม่ให้กระทบการบันทึกเอกสาร
			// log.Printf("Failed to create notification: %v", err)
		}
	}
	
	return c.Status(201).JSON(reportOutline)
}

func GetReportOutlineByTrainingID(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	trainingID := c.Params("trainingId")
	
	var reportOutline models.ReportOutline
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Where("training_id = ?", trainingID).First(&reportOutline).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Report outline not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(reportOutline)
}

func UpdateReportOutline(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	reportOutlineID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var reportOutline models.ReportOutline
	if err := db.First(&reportOutline, reportOutlineID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Report outline not found"})
	}

	var updateData models.ReportOutline
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// อัปเดตข้อมูล
	if err := db.Model(&reportOutline).Updates(&updateData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update report outline"})
	}

	// ดึงข้อมูลที่อัปเดตแล้ว
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&reportOutline, reportOutlineID)
	
	return c.JSON(reportOutline)
}

func DeleteReportOutline(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	reportOutlineID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := db.Delete(&models.ReportOutline{}, reportOutlineID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete report outline"})
	}

	return c.JSON(fiber.Map{"message": "Report outline deleted successfully"})
}

func GetAllReportOutlines(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var reportOutlines []models.ReportOutline
	
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Find(&reportOutlines).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(reportOutlines)
}