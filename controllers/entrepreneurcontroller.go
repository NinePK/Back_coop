package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Entrepreneur (SELECT)
func GetEntrepreneurs(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var entrepreneurs []models.Entrepreneur
	if result := db.Preload("Tambon").Preload("Tambon.Amphur.Province").Find(&entrepreneurs); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(entrepreneurs)
}

// Get Entrepreneur by ID (SELECT)
func GetEntrepreneur(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Entrepreneur models.Entrepreneur
	if result := db.Preload("Tambon").Preload("Tambon.Amphur.Province").Preload("Jobs").
		Joins("JOIN tambon ON tambon_id = tambon.id JOIN amphur ON tambon.amphur_id = amphur.id  JOIN province ON province.id = amphur.province_id ").
		First(&Entrepreneur, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Entrepreneur not found"})
	}
	return c.JSON(Entrepreneur)
}

// Create new Entrepreneur (INSERT)
func CreateEntrepreneur(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Entrepreneur := new(models.Entrepreneur)
	if err := c.BodyParser(Entrepreneur); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON", "detail": err.Error()})
	}

	if result := db.Create(&Entrepreneur); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Entrepreneur)
}

// Update existing Entrepreneur (UPDATE)
func UpdateEntrepreneur(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Entrepreneur models.Entrepreneur
	if result := db.First(&Entrepreneur, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Entrepreneur not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Entrepreneur); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	// log.Println(Entrepreneur)

	var dbSave = db.Model(&Entrepreneur)
	if Entrepreneur.MoobanID == 0 {
		dbSave.Omit("mooban_id")
	}
	if Entrepreneur.TambonID == 0 {
		dbSave.Omit("tambon_id")
	}

	dbSave.Save(&Entrepreneur)
	return c.JSON(Entrepreneur)
}

// Delete Entrepreneur by ID (DELETE)
func DeleteEntrepreneur(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Entrepreneur models.Entrepreneur
	if result := db.First(&Entrepreneur, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Entrepreneur not found"})
	}

	// db.Delete(&Entrepreneur)
	db.First(&Entrepreneur, id).Update("enable", 0)
	return c.JSON(fiber.Map{"message": "Entrepreneur Hide successfully"})
}

func GetEntrepreneurOption(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var entrepreneurs []models.Entrepreneur

	totalEntrepreneurs := int64(0)

	limit := c.QueryInt("perPage", 0)
	if limit <= 0 {
		limit = 10
	}

	offset := c.QueryInt("page", 0)
	if offset != 0 {
		offset = (offset - 1) * limit
	}

	query := db.Preload("Tambon").Preload("Tambon.Amphur.Province").Model(&entrepreneurs).
		Joins("JOIN tambon ON tambon_id = tambon.id JOIN amphur ON tambon.amphur_id = amphur.id  JOIN province ON province.id = amphur.province_id ")

	enable := c.Query("enable", "")
	if enable != "" {
		query = query.Where("enable = ?", enable)
	}

	name := c.Query("name", "")
	if name != "" {
		query = query.Where("name_th LIKE ? OR name_en LIKE ?", "%"+name+"%", "%"+name+"%")
	}

	business := c.Query("business", "")
	if name != "" {
		query = query.Where("business LIKE ? ", "%"+business+"%")
	}

	province_id := c.QueryInt("province")
	if province_id != 0 {
		query = query.Where("amphur.province_id = ?", province_id)
	}

	// start := time.Now()
	totalPages := int64(0)
	count := c.Query("count", "1")
	if count == "1" {
		resultAll := query.Find(&entrepreneurs)
		if resultAll.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": resultAll.Error.Error()})
		}
		resultAll.Count(&totalEntrepreneurs)
		totalPages = totalEntrepreneurs / int64(limit)
	}

	// log.Printf("Count execution time: %s\n", time.Since(start))

	// start = time.Now()

	results := query.Offset(offset).Limit(limit).Find(&entrepreneurs)
	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": results.Error.Error()})
	}

	// log.Printf("Fetch execution time: %s\n", time.Since(start))

	return c.JSON(map[string]interface{}{"entrepreneurs": entrepreneurs, "totalPages": totalPages + 1})
}
