package controllers

import (
	"coop_back/models"
	"log"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Role (SELECT)
func GetRoles(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var roles []models.Role
	if result := db.Find(&roles); result.Error != nil {
		log.Println(result.Error.Error())
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	log.Println(c.JSON(roles))
	return c.JSON(roles)
}

// Get Role by ID (SELECT)
func GetRole(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Role models.Role
	if result := db.First(&Role, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Role not found"})
	}
	return c.JSON(Role)
}

// Create new Role (INSERT)
func CreateRole(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Role := new(models.Role)
	if err := c.BodyParser(Role); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := db.Create(&Role); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Role)
}

// Update existing Role (UPDATE)
func UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Role models.Role
	if result := db.First(&Role, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Role not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Role); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Role)
	return c.JSON(Role)
}

// Delete Role by ID (DELETE)
func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Role models.Role
	if result := db.First(&Role, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Role not found"})
	}

	db.Delete(&Role)
	return c.JSON(fiber.Map{"message": "Role deleted successfully"})
}
