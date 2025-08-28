package controllers

	import (
		"coop_back/models"
	
		"github.com/gofiber/fiber/v2"
	
		"gorm.io/gorm"
	)
	
	// Get all Plan (SELECT)
	func GetPlans(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		var plans []models.Plan
		if result := db.Find(&plans); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(plans)
	}
	
	// Get Plan by ID (SELECT)
	func GetPlan(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
		var Plan models.Plan
		if result := db.First(&Plan, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Plan not found"})
		}
		return c.JSON(Plan)
	}
	
	// Create new Plan (INSERT)
	func CreatePlan(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		Plan := new(models.Plan)
		if err := c.BodyParser(Plan); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		if result := db.Create(&Plan); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(Plan)
	}
	
	// Update existing Plan (UPDATE)
	func UpdatePlan(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var Plan models.Plan
		if result := db.First(&Plan, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Plan not found"})
		}
	
		// Parse request body into struct
		if err := c.BodyParser(&Plan); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		db.Save(&Plan)
		return c.JSON(Plan)
	}
	
	// Delete Plan by ID (DELETE)
	func DeletePlan(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var Plan models.Plan
		if result := db.First(&Plan, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Plan not found"})
		}
	
		db.Delete(&Plan)
		return c.JSON(fiber.Map{"message": "Plan deleted successfully"})
	}
	