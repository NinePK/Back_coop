package controllers

import (
	"strconv"

	"coop_back/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSelfEvaluation(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var selfEvaluation models.SelfEvaluation
	
	if err := c.BodyParser(&selfEvaluation); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// ตรวจสอบว่า training_id มีอยู่จริง
	var training models.Training
	if err := db.First(&training, selfEvaluation.TrainingID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}

	// สร้าง record ใหม่
	if err := db.Create(&selfEvaluation).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create self evaluation"})
	}

	// ดึงข้อมูลพร้อม relation
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&selfEvaluation, selfEvaluation.ID)
	
	// ส่งการแจ้งเตือนไปยังอาจารย์ที่ปรึกษา
	if training.TeacherID1 > 0 {
		if err := CreateDocumentNotification(db, training.UserID, training.TeacherID1, "coop12", "", selfEvaluation.ID); err != nil {
			// Log error แต่ไม่ return error เพื่อไม่ให้กระทบการบันทึกเอกสาร
			// log.Printf("Failed to create notification: %v", err)
		}
	}
	
	return c.Status(201).JSON(selfEvaluation)
}

func GetSelfEvaluationByTrainingID(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	trainingID := c.Params("trainingId")
	
	var selfEvaluation models.SelfEvaluation
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Where("training_id = ?", trainingID).First(&selfEvaluation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Self evaluation not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(selfEvaluation)
}

func UpdateSelfEvaluation(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	selfEvaluationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var selfEvaluation models.SelfEvaluation
	if err := db.First(&selfEvaluation, selfEvaluationID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Self evaluation not found"})
	}

	var updateData models.SelfEvaluation
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// อัปเดตข้อมูล
	if err := db.Model(&selfEvaluation).Updates(&updateData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update self evaluation"})
	}

	// ดึงข้อมูลที่อัปเดตแล้ว
	db.Preload("Training.User").Preload("Training.Job.Entrepreneur").First(&selfEvaluation, selfEvaluationID)
	
	return c.JSON(selfEvaluation)
}

func DeleteSelfEvaluation(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	selfEvaluationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := db.Delete(&models.SelfEvaluation{}, selfEvaluationID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete self evaluation"})
	}

	return c.JSON(fiber.Map{"message": "Self evaluation deleted successfully"})
}

func GetAllSelfEvaluations(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var selfEvaluations []models.SelfEvaluation
	
	if err := db.Preload("Training.User").Preload("Training.Job.Entrepreneur").Find(&selfEvaluations).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(selfEvaluations)
}

// GetSelfEvaluationStatistics - สำหรับดูสถิติการประเมิน
func GetSelfEvaluationStatistics(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var stats struct {
		TotalEvaluations      int64   `json:"totalEvaluations"`
		AverageScore          float64 `json:"averageScore"`
		HighestScore          float64 `json:"highestScore"`
		LowestScore           float64 `json:"lowestScore"`
		ExcellentCount        int64   `json:"excellentCount"`         // >= 8.0
		GoodCount             int64   `json:"goodCount"`              // 6.5-7.99
		SatisfactoryCount     int64   `json:"satisfactoryCount"`      // 5.0-6.49
		NeedsImprovementCount int64   `json:"needsImprovementCount"`  // < 5.0
	}

	// นับจำนวนการประเมินทั้งหมด
	db.Model(&models.SelfEvaluation{}).Count(&stats.TotalEvaluations)

	if stats.TotalEvaluations == 0 {
		return c.JSON(stats)
	}

	// คำนวณคะแนนเฉลี่ย
	var result struct {
		AverageScore float64 `gorm:"column:avg_score"`
		MaxScore     float64 `gorm:"column:max_score"`
		MinScore     float64 `gorm:"column:min_score"`
	}

	db.Model(&models.SelfEvaluation{}).
		Select("AVG(average_score) as avg_score, MAX(average_score) as max_score, MIN(average_score) as min_score").
		Scan(&result)

	stats.AverageScore = result.AverageScore
	stats.HighestScore = result.MaxScore
	stats.LowestScore = result.MinScore

	// นับตามเกรด
	db.Model(&models.SelfEvaluation{}).Where("average_score >= 8.0").Count(&stats.ExcellentCount)
	db.Model(&models.SelfEvaluation{}).Where("average_score >= 6.5 AND average_score < 8.0").Count(&stats.GoodCount)
	db.Model(&models.SelfEvaluation{}).Where("average_score >= 5.0 AND average_score < 6.5").Count(&stats.SatisfactoryCount)
	db.Model(&models.SelfEvaluation{}).Where("average_score < 5.0").Count(&stats.NeedsImprovementCount)

	return c.JSON(stats)
}