package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Job (SELECT)
func GetJobs(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var jobs []models.Job
	if result := db.Preload("Entrepreneur").Find(&jobs); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(jobs)
}

// Get Job by ID (SELECT)
func GetJob(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Job models.Job
	if result := db.Preload("Entrepreneur").First(&Job, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Job not found"})
	}
	return c.JSON(Job)
}

// Create new Job (INSERT)
func CreateJob(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Job := new(models.Job)
	if err := c.BodyParser(Job); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := db.Create(&Job); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Job)
}

// Update existing Job (UPDATE)
func UpdateJob(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Job models.Job
	if result := db.First(&Job, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Job not found"})
	}

	// Parse request body into struct
	if err := c.BodyParser(&Job); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db.Save(&Job)
	return c.JSON(Job)
}

// Delete Job by ID (DELETE)
func DeleteJob(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Job models.Job
	if result := db.First(&Job, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Job not found"})
	}

	db.Delete(&Job)
	return c.JSON(fiber.Map{"message": "Job deleted successfully"})
}

func GetJobOption(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var jobs []models.Job

	totalJobs := int64(0)

	limit := c.QueryInt("perPage", 0)
	if limit <= 0 {
		limit = 10
	}

	offset := c.QueryInt("page", 0)
	if offset != 0 {
		offset = (offset - 1) * limit
	}

	query := db.Preload("Entrepreneur").Model(&jobs).
		Joins("JOIN entrepreneur ON entrepreneur_id = entrepreneur.id")

	name := c.Query("name", "")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	job_des := c.Query("jobDes", "")
	if job_des != "" {
		query = query.Where("job_des LIKE ? ", "%"+job_des+"%")
	}

	// start := time.Now()
	totalPages := int64(0)
	count := c.Query("count", "1")
	if count == "1" {
		resultAll := query.Find(&jobs)
		if resultAll.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": resultAll.Error.Error()})
		}
		resultAll.Count(&totalJobs)
		totalPages = totalJobs / int64(limit)
	}

	// log.Printf("Count execution time: %s\n", time.Since(start))

	// start = time.Now()

	results := query.Offset(offset).Limit(limit).Find(&jobs)
	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": results.Error.Error()})
	}

	// log.Printf("Fetch execution time: %s\n", time.Since(start))

	return c.JSON(map[string]interface{}{"jobs": jobs, "totalPages": totalPages + 1})
}
