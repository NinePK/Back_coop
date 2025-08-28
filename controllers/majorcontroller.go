package controllers

import (
	"coop_back/models"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get all Majors (SELECT)
func GetMajors(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var Majors []models.Major
	if result := db.Find(&Majors); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Majors)
}

// Get Major by ID (SELECT)
func GetMajor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	// log.Infof("major id: %s", id)
	var Major models.Major
	if result := db.First(&Major, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Major not found"})
	}
	return c.JSON(Major)
}

func GetMajorWhere(c *fiber.Ctx) error {
	key, _ := url.QueryUnescape(c.Params("param"))
	val, _ := url.QueryUnescape(c.Params("val"))
	// log.Infof("key: %s, val: %s", key, val)
	db := c.Locals("db").(*gorm.DB)
	var Major models.Major
	if result := db.Where(key+" = ?", val).First(&Major); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Major not found"})
	}
	return c.JSON(Major)
}

// Create new Major (INSERT)
func CreateMajor(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	type Request struct {
		MajorTh   string `json:"majorTh"`
		MajorEn   string `json:"majorEn"`
		FacultyTh string `json:"FacultyTh"`
		FacultyEn string `json:"FacultyEn"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Step 1: Check if the majorTh already exists
	var major models.Major
	if err := db.Where("major_th = ?", request.MajorTh).First(&major).Error; err == nil {
		// Major already exists, return its ID
		return c.JSON(fiber.Map{
			"major_id": major.ID,
		})
	}

	// Step 2: Check if the FacultyTh already exists
	var faculty models.Faculty
	if err := db.Where("faculty_th = ?", request.FacultyTh).First(&faculty).Error; err != nil {
		// Faculty does not exist, create a new one
		faculty = models.Faculty{
			FacultyTh: request.FacultyTh,
			FacultyEn: request.FacultyEn,
		}
		if err := db.Create(&faculty).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create faculty",
			})
		}
	}

	// Step 3: Create new major
	newMajor := models.Major{
		MajorTh:   request.MajorTh,
		MajorEn:   request.MajorEn,
		FacultyID: faculty.ID,
	}

	if err := db.Create(&newMajor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create major",
		})
	}

	// Step 4: Return the new major ID
	return c.JSON(fiber.Map{
		"major_id": newMajor.ID,
	})
}

// Update existing Major (UPDATE)
func UpdateMajor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Major models.Major
	if result := db.First(&Major, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Major not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Major); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Major)
	return c.JSON(Major)
}

// Delete Major by ID (DELETE)
func DeleteMajor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Major models.Major
	if result := db.First(&Major, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Major not found"})
	}

	db.Delete(&Major)
	return c.JSON(fiber.Map{"message": "Major deleted successfully"})
}
