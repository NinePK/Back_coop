package controllers

import (
	"coop_back/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all Trainings (SELECT)
func GetTrainings(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var Trainings []models.Training
	if result := db.Preload("Role").Find(&Trainings); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Trainings)
}

// Get Training by ID (SELECT)
func GetTraining(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Training models.Training
	if result := db.Preload("Role").First(&Training, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}
	return c.JSON(Training)
}

// Create new Training (INSERT)
func CreateTraining(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	Training := new(models.Training)

	// TrainingUpdate := models.Training{
	// 	StartDate: time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), // Set date with time zeroed
	// }

	// Training.StartDate = TrainingUpdate.StartDate

	if err := c.BodyParser(Training); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error(), "Training": c.JSON(Training)})
	}

	if result := db.Preload("Role").Create(&Training); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Training)
}

// Update existing Training (UPDATE)
func UpdateTraining(c *fiber.Ctx) error {
	// id := c.Body["id"]

	db := c.Locals("db").(*gorm.DB)

	Training := new(models.Training)

	TrainingParser := new(models.Training)

	if err := c.BodyParser(TrainingParser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error updateTraining(Map Form) ": err.Error(), "Training": c.JSON(Training)})
	}

	if result := db.First(&Training, TrainingParser.ID); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error updateTraining(Search)": "Training not found"})
	}

	if err := c.BodyParser(Training); err != nil {
		return c.Status(400).JSON(fiber.Map{"error updateTraining(Map Form) ": err.Error(), "Training": c.JSON(Training)})
	}

	var dbSave = db.Model(&Training)

	if Training.MoobanID == 0 {
		dbSave.Omit("moobanId")
	}
	if Training.TambonID == 0 {
		dbSave.Omit("tambonId")
	}
	if Training.InchargeID1 == 0 {
		dbSave.Omit("inchargeId1")
	}
	if Training.InchargeID2 == 0 {
		dbSave.Omit("inchargeId2")
	}
	if Training.TeacherID1 == 0 {
		dbSave.Omit("teacherId1")
	}
	if Training.TeacherID2 == 0 {
		dbSave.Omit("teacherId2")
	}

	dbSave.Debug().Updates(&Training)
	return c.JSON(Training)

}

// Delete Training by ID (DELETE)
func DeleteTraining(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var Training models.Training
	if result := db.First(&Training, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}

	db.Delete(&Training)
	return c.JSON(fiber.Map{"message": "Training deleted successfully"})
}

func GetTrainingsByUser(c *fiber.Ctx) error {

	user_id, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be an integer")
	}

	semester_id, err := c.ParamsInt("semester_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("semester must be an integer")
	}

	db := c.Locals("db").(*gorm.DB)
	// Training := new(models.Training)
	var Trainings []models.Training

	// log.Println(c.JSON(Training))

	if result := db.Debug().Preload("Job").Preload("Job.Entrepreneur").Preload("Tambon").Preload("Tambon.Amphur").
		Where("user_id = ? and semester_id = ? ", user_id, semester_id).Find(&Trainings); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}
	return c.JSON(Trainings)
}
