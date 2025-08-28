package controllers

import (
	"strconv"

	"coop_back/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateJobDetails(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var jobDetails models.JobDetails
	
	if err := c.BodyParser(&jobDetails); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// ตรวจสอบว่า training_id มีอยู่จริง
	var training models.Training
	if err := db.First(&training, jobDetails.TrainingID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}

	// สร้าง record ใหม่
	if err := db.Create(&jobDetails).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create job details"})
	}

	// ดึงข้อมูลพร้อม relation
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&jobDetails, jobDetails.ID)
	
	// ส่งการแจ้งเตือนไปยังอาจารย์ที่ปรึกษา
	if training.TeacherID1 > 0 {
		if err := CreateDocumentNotification(db, training.UserID, training.TeacherID1, "coop11", "", jobDetails.ID); err != nil {
			// Log error แต่ไม่ return error เพื่อไม่ให้กระทบการบันทึกเอกสาร
			// log.Printf("Failed to create notification: %v", err)
		}
	}
	
	return c.Status(201).JSON(jobDetails)
}

func GetJobDetailsByTrainingID(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	trainingID := c.Params("trainingId")
	
	var jobDetails models.JobDetails
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Where("training_id = ?", trainingID).First(&jobDetails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Job details not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(jobDetails)
}

func UpdateJobDetails(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	jobDetailsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var jobDetails models.JobDetails
	if err := db.First(&jobDetails, jobDetailsID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Job details not found"})
	}

	var updateData models.JobDetails
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// อัปเดตข้อมูล
	if err := db.Model(&jobDetails).Updates(&updateData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update job details"})
	}

	// ดึงข้อมูลที่อัปเดตแล้ว
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&jobDetails, jobDetailsID)
	
	return c.JSON(jobDetails)
}

func DeleteJobDetails(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	jobDetailsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := db.Delete(&models.JobDetails{}, jobDetailsID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete job details"})
	}

	return c.JSON(fiber.Map{"message": "Job details deleted successfully"})
}

func GetAllJobDetails(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var jobDetails []models.JobDetails
	
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Find(&jobDetails).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(jobDetails)
}