package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Semester (SELECT)
func GetSemesters(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var semesters []models.Semester
	if result := db.Find(&semesters); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(semesters)
}

// Get Semester by ID (SELECT)
func GetSemester(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Semester models.Semester
	if result := db.First(&Semester, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Semester not found"})
	}
	return c.JSON(Semester)
}

func GetCurrentSemester(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var Semester models.Semester
	if result := db.First(&Semester, "is_current = 1 "); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Current Semester not defined yet"})
	}
	return c.JSON(Semester)
}

// Create new Semester (INSERT)
func CreateSemester(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Semester := new(models.Semester)
	if err := c.BodyParser(Semester); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := db.Create(&Semester); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Semester)
}

// Update existing Semester (UPDATE)
func UpdateSemester(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Semester models.Semester
	if result := db.First(&Semester, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Semester not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Semester); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Semester)
	return c.JSON(Semester)
}

func UpdateSemesterByData(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var Semester models.Semester
	if err := c.BodyParser(&Semester); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON", "data": c.JSON(Semester)})
	}

	if result := db.First(&Semester, "semester = ? AND year = ?", Semester.Semester, Semester.Year); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Semester not found"})
	}

	Semester.IsCurrent = 1

	db.Save(&Semester)
	return c.JSON(Semester)
}

// Delete Semester by ID (DELETE)
func DeleteSemester(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Semester models.Semester
	if result := db.First(&Semester, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Semester not found"})
	}

	db.Delete(&Semester)
	return c.JSON(fiber.Map{"message": "Semester deleted successfully"})
}
