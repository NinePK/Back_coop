package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Province (SELECT)
func GetProvinces(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var provinces []models.Province
	if result := db.Order("value").Find(&provinces); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(provinces)
}

// Get Province by ID (SELECT)
func GetProvince(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Province models.Province
	if result := db.First(&Province, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Province not found"})
	}
	return c.JSON(Province)
}

// Create new Province (INSERT)
func CreateProvince(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Province := new(models.Province)
	if err := c.BodyParser(Province); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := db.Create(&Province); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Province)
}

// Update existing Province (UPDATE)
func UpdateProvince(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Province models.Province
	if result := db.First(&Province, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Province not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Province); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Province)
	return c.JSON(Province)
}

// Delete Province by ID (DELETE)
func DeleteProvince(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Province models.Province
	if result := db.First(&Province, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Province not found"})
	}

	db.Delete(&Province)
	return c.JSON(fiber.Map{"message": "Province deleted successfully"})
}
