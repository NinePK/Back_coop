package controllers

import (
	"coop_back/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get all Incharge (SELECT)
func GetIncharges(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var incharges []models.Incharge
	if result := db.Preload("Entrepreneur").Find(&incharges); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(incharges)
}

// Get Incharge by ID (SELECT)
func GetIncharge(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Incharge models.Incharge
	if result := db.Preload("Entrepreneur").First(&Incharge, id); result.Error != nil {
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

	// Hash password if provided
	if Incharge.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Incharge.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		Incharge.Password = string(hashedPassword)
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

	// Hash password if provided and changed
	if Incharge.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Incharge.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		Incharge.Password = string(hashedPassword)
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

// AuthenticateIncharge - ตรวจสอบการล็อกอินของพนักงานที่ปรึกษา
func AuthenticateIncharge(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	
	// ค้นหาพนักงานที่ปรึกษาด้วยอีเมล (ใช้ username field เป็น email)
	var incharge models.Incharge
	if err := db.Preload("Entrepreneur").Where("username = ?", request.Email).First(&incharge).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
		})
	}
	
	// ตรวจสอบรหัสผ่าน
	if incharge.Password != "" {
		// ถ้ามีการ hash รหัสผ่าน
		err := bcrypt.CompareHashAndPassword([]byte(incharge.Password), []byte(request.Password))
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
			})
		}
	} else {
		// ถ้าไม่มีรหัสผ่านในระบบ ให้ส่งข้อความแจ้ง
		return c.Status(401).JSON(fiber.Map{
			"error": "ยังไม่ได้ตั้งรหัสผ่าน กรุณาติดต่อผู้ดูแลระบบ",
		})
	}
	
	// ส่งข้อมูลพนักงานที่ปรึกษากลับไป
	return c.JSON(fiber.Map{
		"id":             incharge.ID,
		"fname":          incharge.Fname,
		"sname":          incharge.Sname,
		"position":       incharge.Position,
		"username":       incharge.Username,
		"entrepreneur_id": incharge.EntrepreneurID,
		"entrepreneur":   incharge.Entrepreneur,
	})
}