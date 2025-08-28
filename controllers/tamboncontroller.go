package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Tambon (SELECT)
func GetTambons(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var tambons []models.Tambon
	
	// ตรวจสอบ query parameter amphur_id
	amphurId := c.Query("amphur_id")
	if amphurId != "" {
		// ถ้ามี amphur_id ให้ filter ตาม amphur_id
		if result := db.Order("value").Find(&tambons, "amphur_id = ?", amphurId); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Tambon not found"})
		}
	} else {
		// ถ้าไม่มี amphur_id ให้ดึงทั้งหมด
		if result := db.Find(&tambons); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
	}
	return c.JSON(tambons)
}

// Get Tambon by ID (SELECT)
func GetTambon(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Tambon models.Tambon
	if result := db.First(&Tambon, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tambon not found"})
	}
	return c.JSON(Tambon)
}

func GetTambonByAmphurId(c *fiber.Ctx) error {
	amphur_id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Tambon []models.Tambon
	if result := db.Order("value").Find(&Tambon, "amphur_id = ?", amphur_id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tambon not found"})
	}
	return c.JSON(Tambon)
}

// Create new Tambon (INSERT)
func CreateTambon(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Tambon := new(models.Tambon)
	if err := c.BodyParser(Tambon); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := db.Create(&Tambon); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Tambon)
}

// Update existing Tambon (UPDATE)
func UpdateTambon(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Tambon models.Tambon
	if result := db.First(&Tambon, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tambon not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Tambon); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Tambon)
	return c.JSON(Tambon)
}

// Delete Tambon by ID (DELETE)
func DeleteTambon(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Tambon models.Tambon
	if result := db.First(&Tambon, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tambon not found"})
	}

	db.Delete(&Tambon)
	return c.JSON(fiber.Map{"message": "Tambon deleted successfully"})
}
