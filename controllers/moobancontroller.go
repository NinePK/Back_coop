package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Mooban (SELECT)
func GetMoobans(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var moobans []models.Mooban
	if result := db.Find(&moobans); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(moobans)
}

// Get Mooban by ID (SELECT)
func GetMooban(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Mooban models.Mooban
	if result := db.First(&Mooban, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mooban not found"})
	}
	return c.JSON(Mooban)
}

func GetMoonbanByTambonId(c *fiber.Ctx) error {
	tambon_id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Mooban []models.Mooban
	if result := db.Order("value").Find(&Mooban, "tambon_id = ?", tambon_id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mooban not found"})
	}
	return c.JSON(Mooban)
}

// Create new Mooban (INSERT)
func CreateMooban(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Mooban := new(models.Mooban)
	if err := c.BodyParser(Mooban); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := db.Create(&Mooban); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Mooban)
}

// Update existing Mooban (UPDATE)
func UpdateMooban(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Mooban models.Mooban
	if result := db.First(&Mooban, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mooban not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Mooban); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Mooban)
	return c.JSON(Mooban)
}

// Delete Mooban by ID (DELETE)
func DeleteMooban(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Mooban models.Mooban
	if result := db.First(&Mooban, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mooban not found"})
	}

	db.Delete(&Mooban)
	return c.JSON(fiber.Map{"message": "Mooban deleted successfully"})
}
