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
	if result := db.Debug().
		Preload("User").
		Preload("User.Major").
		Preload("User.Major.Faculty").
		Preload("Job").
		Preload("Job.Entrepreneur").
		Preload("Semester").
		Preload("Mooban").
		Preload("Tambon").
		Preload("Tambon.Amphur").
		Preload("Tambon.Amphur.Province").
		Preload("Teacher1").
		Preload("Teacher2").
		Preload("Incharge1").
		Preload("Incharge2").
		Find(&Trainings); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(Trainings)
}

// Get Training by ID (SELECT)
func GetTraining(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var Training models.Training
	if result := db.Debug().
		Preload("User").
		Preload("User.Major").
		Preload("User.Major.Faculty").
		Preload("Job").
		Preload("Job.Entrepreneur").
		Preload("Semester").
		Preload("Mooban").
		Preload("Tambon").
		Preload("Tambon.Amphur").
		Preload("Tambon.Amphur.Province").
		Preload("Teacher1").
		Preload("Teacher2").
		Preload("Incharge1").
		Preload("Incharge2").
		First(&Training, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Training not found"})
	}
	return c.JSON(Training)
}

// Create new Training (INSERT)
func CreateTraining(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	// Parse the request into the custom struct first
	var trainingRequest models.TrainingRequest
	if err := c.BodyParser(&trainingRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// ดึงข้อมูล Job และ Entrepreneur ก่อน
	var job models.Job
	if err := db.Preload("Entrepreneur").First(&job, trainingRequest.JobID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Job not found"})
	}

	// Convert the request to the Training model
	Training := &models.Training{
		UserID:     trainingRequest.UserID,
		JobID:      trainingRequest.JobID,
		SemesterID: trainingRequest.SemesterID,
		StartDate:  &trainingRequest.StartDate,
		EndDate:    &trainingRequest.EndDate,
		Coop:       trainingRequest.Coop,
		Status:     trainingRequest.Status,
		// ดึงข้อมูลจาก Entrepreneur
		Address:        job.Entrepreneur.Address,
		Tel:            job.Entrepreneur.Tel,
		Email:          job.Entrepreneur.Email,
		NameMentor:     job.Entrepreneur.Manager,
		PositionMentor: job.Entrepreneur.ManagerPosition,
		DeptMentor:     job.Entrepreneur.ManagerDept,
		TelMentor:      job.Entrepreneur.ContactTel,
		EmailMentor:    job.Entrepreneur.ContactEmail,
		// ดึงข้อมูลจาก Job
		JobPosition:    job.Name,
		JobDes:         job.JobDes,
	}

	// Handle remaining nullable fields from request (override auto-filled data if provided)
	if trainingRequest.Address != nil && *trainingRequest.Address != "" {
		Training.Address = *trainingRequest.Address
	}
	if trainingRequest.MoobanID != nil {
		Training.MoobanID = *trainingRequest.MoobanID
	}
	if trainingRequest.TambonID != nil {
		Training.TambonID = *trainingRequest.TambonID
	}
	if trainingRequest.Tel != nil && *trainingRequest.Tel != "" {
		Training.Tel = *trainingRequest.Tel
	}
	if trainingRequest.Email != nil && *trainingRequest.Email != "" {
		Training.Email = *trainingRequest.Email
	}
	if trainingRequest.Lat != nil {
		Training.Lat = *trainingRequest.Lat
	}
	if trainingRequest.Long != nil {
		Training.Long = *trainingRequest.Long
	}
	if trainingRequest.TeacherID1 != nil {
		Training.TeacherID1 = *trainingRequest.TeacherID1
	}
	if trainingRequest.TeacherID2 != nil {
		Training.TeacherID2 = *trainingRequest.TeacherID2
	}
	if trainingRequest.InchargeID1 != nil {
		Training.InchargeID1 = *trainingRequest.InchargeID1
	}
	if trainingRequest.InchargeID2 != nil {
		Training.InchargeID2 = *trainingRequest.InchargeID2
	}

	if result := db.Create(&Training); result.Error != nil {
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
