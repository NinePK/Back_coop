package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Amphur (SELECT)
func GetAmphurs(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var amphurs []models.Amphur
	if result := db.Find(&amphurs); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(amphurs)
}

// Get Amphur by ID (SELECT)
func GetAmphur(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Amphur models.Amphur
	if result := db.First(&Amphur, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Amphur not found"})
	}
	return c.JSON(Amphur)
}

func GetAmphurByProvinceId(c *fiber.Ctx) error {

	// origin := c.Get("Origin")
	// log.Println("Origin:", origin)

	province_id := c.Params("id")
	// log.Printf("province_id: %s", province_id)

	db := c.Locals("db").(*gorm.DB)
	var Amphur []models.Amphur
	if result := db.Order("value").Find(&Amphur, "province_id = ?", province_id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Amphur not found"})
	}
	return c.JSON(Amphur)
}

// Create new Amphur (INSERT)
func CreateAmphur(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Amphur := new(models.Amphur)
	if err := c.BodyParser(Amphur); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := db.Create(&Amphur); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Amphur)
}

// Update existing Amphur (UPDATE)
func UpdateAmphur(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Amphur models.Amphur
	if result := db.First(&Amphur, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Amphur not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Amphur); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Amphur)
	return c.JSON(Amphur)
}

// Delete Amphur by ID (DELETE)
func DeleteAmphur(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Amphur models.Amphur
	if result := db.First(&Amphur, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Amphur not found"})
	}

	db.Delete(&Amphur)
	return c.JSON(fiber.Map{"message": "Amphur deleted successfully"})
}
