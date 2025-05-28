package controllers

import (
	"coop_back/models"
	"log"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// Get all users (SELECT)
func GetUsers(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var users []models.User
	if result := db.Preload("Role").Find(&users); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(users)
}

func GetUsersOption(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	// type User struct {
	// 	ID      uint   `json:"id"`
	// 	Fname   string `json:"name"`
	// 	Sname   string `json:"email"`
	// 	Major   string `json:"major"`
	// 	Role    string `json:"role"`
	// 	Faculty string `json:"faculty"`
	// }

	// var params SearchParams

	var users []models.User
	// var user models.User
	totalUsers := int64(0)

	limit := c.QueryInt("perPage", 0)
	if limit <= 0 {
		limit = 10
	}

	offset := c.QueryInt("page", 0)
	if offset != 0 {
		offset = (offset - 1) * limit
	}

	query := db.Debug().Preload("Major").Preload("Major.Faculty").Preload("Role").Model(&users).
		Joins("JOIN major ON major.id = user.major_id JOIN  role ON role.id = user.role_id JOIN  faculty ON major.faculty_id = faculty.id")

	fname := c.Query("fname", "")
	if fname != "" {
		log.Println("fname: ", fname)
		query = query.Where("fname LIKE ?", "%"+fname+"%")
	}

	sname := c.Query("sname", "")
	if sname != "" {
		query = query.Where("sname LIKE ?", "%"+sname+"%")
	}

	major := c.QueryInt("major")
	if major != 0 {
		query = query.Where("major_id = ?", major)
	}

	faculty := c.QueryInt("faculty")
	if faculty != 0 {
		query = query.Where("major.faculty_id = ?", faculty)
	}

	role := c.QueryInt("role")
	if role != 0 {
		query = query.Where("role_id = ?", role)
	}

	log.Println("query: ", limit, offset, faculty, major, role)

	// start := time.Now()

	resultAll := query.Find(&users)
	if resultAll.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": resultAll.Error.Error()})
	}
	// totalPages := resultAll.RowsAffected
	resultAll.Count(&totalUsers)
	totalPages := totalUsers / int64(limit)

	// log.Printf("Count execution time: %s\n", time.Since(start))

	// start = time.Now()

	resultAll = query.Offset(offset).Limit(limit).Find(&users)
	if resultAll.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": resultAll.Error.Error()})
	}

	// log.Printf("Fetch execution time: %s\n", time.Since(start))

	return c.JSON(map[string]interface{}{"users": users, "totalPages": totalPages + 1})
}

// Get user by ID (SELECT)
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)
	var user models.User
	if result := db.Preload("Role").First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// Create new user (INSERT)
func CreateUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error(), "user": c.JSON(user)})
	}

	if result := db.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	newUser := new(models.User)
	if resultSelect := db.Preload("Role").Preload("Major").Preload("Major.Faculty").First(&newUser, user.ID); resultSelect.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// log.Println(newUser)

	return c.JSON(newUser)
}

// Update existing user (UPDATE)
func UpdateUser(c *fiber.Ctx) error {

	db := c.Locals("db").(*gorm.DB)

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error updateUser(Map Form) ": err.Error(), "user": c.JSON(user)})
	}

	if result := db.First(&user, user.ID); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error updateUser(Search)": "User not found"})
	}

	db.Save(&user)
	return c.JSON(user)
}

// Delete user by ID (DELETE)
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var user models.User
	if result := db.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	db.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

func SearchUser(c *fiber.Ctx) error {

	db := c.Locals("db").(*gorm.DB)
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	log.Println(c.JSON(user))

	if result := db.Preload("Major").Preload("Major.Faculty").Preload("Role").Where("fname = ? AND sname = ?", user.Fname, user.Sname).First(&user); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}
func SearchUserByUsername(c *fiber.Ctx) error {
	// ดึงข้อมูลจาก body
	db := c.Locals("db").(*gorm.DB)
	user := new(models.User)

	username := c.Params("username") // ใช้ username ที่ส่งมาจาก URL

	// ค้นหาผู้ใช้จาก username
	if result := db.Preload("Major").Preload("Major.Faculty").Preload("Role").Where("username = ?", username).First(&user); result.Error != nil {
		// หากไม่พบผู้ใช้
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// ส่งข้อมูลผู้ใช้กลับไป
	return c.JSON(user)
}
