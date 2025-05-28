package controllers

	import (
		"coop_back/models"
	
		"github.com/gofiber/fiber/v2"
	
		"gorm.io/gorm"
	)
	
	// Get all Faculty (SELECT)
	func GetFacultys(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		var facultys []models.Faculty
		if result := db.Find(&facultys); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(facultys)
	}
	
	// Get Faculty by ID (SELECT)
	func GetFaculty(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
		var Faculty models.Faculty
		if result := db.First(&Faculty, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Faculty not found"})
		}
		return c.JSON(Faculty)
	}
	
	// Create new Faculty (INSERT)
	func CreateFaculty(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		Faculty := new(models.Faculty)
		if err := c.BodyParser(Faculty); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		if result := db.Create(&Faculty); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(Faculty)
	}
	
	// Update existing Faculty (UPDATE)
	func UpdateFaculty(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var Faculty models.Faculty
		if result := db.First(&Faculty, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Faculty not found"})
		}
	
		// Parse request body into struct
		if err := c.BodyParser(&Faculty); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		db.Save(&Faculty)
		return c.JSON(Faculty)
	}
	
	// Delete Faculty by ID (DELETE)
	func DeleteFaculty(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var Faculty models.Faculty
		if result := db.First(&Faculty, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Faculty not found"})
		}
	
		db.Delete(&Faculty)
		return c.JSON(fiber.Map{"message": "Faculty deleted successfully"})
	}
	