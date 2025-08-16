package controllers

import (
	"coop_back/models"
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Coop04Request struct {
	Document struct {
		UserID       int64  `json:"user_id"`
		TrainingID   int64  `json:"training_id"`
		DocumentType string `json:"document_type"`
		DocumentName string `json:"document_name"`
		Status       string `json:"status"`
	} `json:"document"`
	Accommodation struct {
		AccommodationType string  `json:"accommodation_type"`
		AccommodationName string  `json:"accommodation_name"`
		RoomNumber        string  `json:"room_number"`
		Address           string  `json:"address"`
		Subdistrict       string  `json:"subdistrict"`
		District          string  `json:"district"`
		Province          string  `json:"province"`
		PostalCode        string  `json:"postal_code"`
		PhoneNumber       string  `json:"phone_number"`
		EmergencyContact  string  `json:"emergency_contact"`
		EmergencyPhone    string  `json:"emergency_phone"`
		EmergencyRelation string  `json:"emergency_relation"`
		TravelMethod      string  `json:"travel_method"`
		TravelDetails     string  `json:"travel_details"`
		DistanceKm        float64 `json:"distance_km"`
		TravelTime        int64   `json:"travel_time"`
	} `json:"accommodation"`
}

func CreateCoop04Accommodation(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var req Coop04Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	// Create accommodation record
	accommodation := models.Coop04Accommodation{
		TrainingID:          req.Document.TrainingID,
		UserID:              req.Document.UserID,
		AccommodationType:   req.Accommodation.AccommodationType,
		AccommodationName:   req.Accommodation.AccommodationName,
		RoomNumber:          req.Accommodation.RoomNumber,
		Address:             req.Accommodation.Address,
		Subdistrict:         req.Accommodation.Subdistrict,
		District:            req.Accommodation.District,
		Province:            req.Accommodation.Province,
		PostalCode:          req.Accommodation.PostalCode,
		PhoneNumber:         req.Accommodation.PhoneNumber,
		EmergencyContact:    req.Accommodation.EmergencyContact,
		EmergencyPhone:      req.Accommodation.EmergencyPhone,
		EmergencyRelation:   req.Accommodation.EmergencyRelation,
		TravelMethod:        req.Accommodation.TravelMethod,
		TravelDetails:       req.Accommodation.TravelDetails,
		DistanceKm:          req.Accommodation.DistanceKm,
		TravelTime:          req.Accommodation.TravelTime,
		Status:              "submitted",
	}

	result := db.Create(&accommodation)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error: " + result.Error.Error(),
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"success": true,
		"message": "บันทึกข้อมูลที่พักเรียบร้อยแล้ว",
		"data": accommodation,
	})
}

func GetCoop04Accommodation(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	userID := c.Query("userId")
	trainingID := c.Query("trainingId")

	var accommodation models.Coop04Accommodation

	if trainingID != "" {
		result := db.Where("training_id = ?", trainingID).Order("created_at DESC").First(&accommodation)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return c.JSON(fiber.Map{"data": nil})
			}
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
	} else if userID != "" {
		result := db.Table("coop04_accommodation ca").
			Joins("JOIN training t ON ca.training_id = t.id").
			Where("t.user_id = ?", userID).
			Order("ca.created_at DESC").
			First(&accommodation)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return c.JSON(fiber.Map{"data": nil})
			}
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
	} else {
		return c.Status(400).JSON(fiber.Map{"error": "userId or trainingId parameter required"})
	}

	return c.JSON(fiber.Map{"data": accommodation})
}

func UpdateCoop04Accommodation(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var req struct {
		ID            int64 `json:"id"`
		Accommodation struct {
			AccommodationType string  `json:"accommodation_type"`
			AccommodationName string  `json:"accommodation_name"`
			RoomNumber        string  `json:"room_number"`
			Address           string  `json:"address"`
			Subdistrict       string  `json:"subdistrict"`
			District          string  `json:"district"`
			Province          string  `json:"province"`
			PostalCode        string  `json:"postal_code"`
			PhoneNumber       string  `json:"phone_number"`
			EmergencyContact  string  `json:"emergency_contact"`
			EmergencyPhone    string  `json:"emergency_phone"`
			EmergencyRelation string  `json:"emergency_relation"`
			TravelMethod      string  `json:"travel_method"`
			TravelDetails     string  `json:"travel_details"`
			DistanceKm        float64 `json:"distance_km"`
			TravelTime        int64   `json:"travel_time"`
		} `json:"accommodation"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	// Get ID from either URL params or request body
	var id int64
	if c.Params("id") != "" {
		idParam, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID in URL"})
		}
		id = int64(idParam)
	} else if req.ID != 0 {
		id = req.ID
	} else {
		return c.Status(400).JSON(fiber.Map{"error": "ID required"})
	}

	var accommodation models.Coop04Accommodation
	result := db.First(&accommodation, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "Accommodation not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Update fields
	accommodation.AccommodationType = req.Accommodation.AccommodationType
	accommodation.AccommodationName = req.Accommodation.AccommodationName
	accommodation.RoomNumber = req.Accommodation.RoomNumber
	accommodation.Address = req.Accommodation.Address
	accommodation.Subdistrict = req.Accommodation.Subdistrict
	accommodation.District = req.Accommodation.District
	accommodation.Province = req.Accommodation.Province
	accommodation.PostalCode = req.Accommodation.PostalCode
	accommodation.PhoneNumber = req.Accommodation.PhoneNumber
	accommodation.EmergencyContact = req.Accommodation.EmergencyContact
	accommodation.EmergencyPhone = req.Accommodation.EmergencyPhone
	accommodation.EmergencyRelation = req.Accommodation.EmergencyRelation
	accommodation.TravelMethod = req.Accommodation.TravelMethod
	accommodation.TravelDetails = req.Accommodation.TravelDetails
	accommodation.DistanceKm = req.Accommodation.DistanceKm
	accommodation.TravelTime = req.Accommodation.TravelTime
	accommodation.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	result = db.Save(&accommodation)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "อัปเดตข้อมูลที่พักเรียบร้อยแล้ว",
		"data": accommodation,
	})
}