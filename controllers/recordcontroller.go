package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Records (SELECT)
func GetRecords(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var Records []models.Record
	if result := db.Preload("Role").Find(&Records); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Records)
}

// Get Record by ID (SELECT)
func GetRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Record models.Record
	if result := db.Preload("Role").First(&Record, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Record not found"})
	}
	return c.JSON(Record)
}

// Create new Record (INSERT)
func CreateRecord(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Record := new(models.Record)
	if err := c.BodyParser(Record); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error(), "Record": c.JSON(Record)})
	}

	if result := db.Preload("Role").Create(&Record); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Record)
}

// Update existing Record (UPDATE)
func UpdateRecord(c *fiber.Ctx) error {

	db := c.Locals("db").(*gorm.DB)

	Record := new(models.Record)

	if err := c.BodyParser(Record); err != nil {
		return c.Status(400).JSON(fiber.Map{"error updateRecord(Map Form) ": err.Error(), "Record": c.JSON(Record)})
	}

	if result := db.First(&Record, Record.ID); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error updateRecord(Search)": "Record not found"})
	}

	db.Save(&Record)
	return c.JSON(Record)
}

// Delete Record by ID (DELETE)
func DeleteRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Record models.Record
	if result := db.First(&Record, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Record not found"})
	}

	db.Delete(&Record)
	return c.JSON(fiber.Map{"message": "Record deleted successfully"})
}

// func SearchRecord(c *fiber.Ctx) error {

// 	db := c.Locals("db").(*gorm.DB)
// 	Record := new(models.Record)
// 	if err := c.BodyParser(&Record); err != nil {
// 		return err
// 	}

// 	log.Println(c.JSON(Record))

// 	if result := db.Where("fname = ? AND sname = ?", Record.Fname, Record.Sname).First(&Record); result.Error != nil {
// 		return c.Status(404).JSON(fiber.Map{"error": "Record not found"})
// 	}
// 	return c.JSON(Record)
// }
