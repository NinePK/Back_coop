package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Weekly reports (SELECT)
func GetWeeklys(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var weeklys []models.Weekly
	
	// Check if training_id is provided
	trainingID := c.Query("training_id")
	if trainingID != "" {
		if result := db.Where("training_id = ?", trainingID).Find(&weeklys); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
	} else {
		if result := db.Find(&weeklys); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
	}
	
	return c.JSON(weeklys)
}

// Get Weekly by ID (SELECT)
func GetWeekly(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var weekly models.Weekly
	if result := db.First(&weekly, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Weekly report not found"})
	}
	return c.JSON(weekly)
}

// Create new Weekly report (INSERT)
func CreateWeekly(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	
	weekly := new(models.Weekly)
	
	if err := c.BodyParser(weekly); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	
	if result := db.Create(&weekly); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	
	return c.JSON(weekly)
}

// Update existing Weekly report (UPDATE)
func UpdateWeekly(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	
	weekly := new(models.Weekly)
	weeklyParser := new(models.Weekly)
	
	if err := c.BodyParser(weeklyParser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	
	if result := db.First(&weekly, weeklyParser.ID); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Weekly report not found"})
	}
	
	if err := c.BodyParser(weekly); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	
	db.Save(&weekly)
	return c.JSON(weekly)
}

// Delete Weekly report by ID (DELETE)
func DeleteWeekly(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	
	var weekly models.Weekly
	if result := db.First(&weekly, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Weekly report not found"})
	}
	
	db.Delete(&weekly)
	return c.JSON(fiber.Map{"message": "Weekly report deleted successfully"})
}

// Get Weekly reports by Training ID
func GetWeeklysByTraining(c *fiber.Ctx) error {
	trainingID := c.Params("training_id")
	db := c.Locals("db").(*gorm.DB)
	var weeklys []models.Weekly
	
	if result := db.Where("training_id = ?", trainingID).Order("week ASC").Find(&weeklys); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	
	return c.JSON(weeklys)
}