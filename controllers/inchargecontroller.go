package controllers

	import (
		"coop_back/models"
	
		"github.com/gofiber/fiber/v2"
	
		"gorm.io/gorm"
	)
	
	// Get all Incharge (SELECT)
	func GetIncharges(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		var incharges []models.Incharge
		if result := db.Find(&incharges); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(incharges)
	}
	
	// Get Incharge by ID (SELECT)
	func GetIncharge(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
		var Incharge models.Incharge
		if result := db.First(&Incharge, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Incharge not found"})
		}
		return c.JSON(Incharge)
	}
	
	// Create new Incharge (INSERT)
	func CreateIncharge(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		Incharge := new(models.Incharge)
		if err := c.BodyParser(Incharge); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		if result := db.Create(&Incharge); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON(Incharge)
	}
	
	// Update existing Incharge (UPDATE)
	func UpdateIncharge(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var Incharge models.Incharge
		if result := db.First(&Incharge, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Incharge not found"})
		}
	
		// Parse request body into struct
		if err := c.BodyParser(&Incharge); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		db.Save(&Incharge)
		return c.JSON(Incharge)
	}
	
	// Delete Incharge by ID (DELETE)
	func DeleteIncharge(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var Incharge models.Incharge
		if result := db.First(&Incharge, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Incharge not found"})
		}
	
		db.Delete(&Incharge)
		return c.JSON(fiber.Map{"message": "Incharge deleted successfully"})
	}
	