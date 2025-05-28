package main

import (
	"log"
	"os"
	"strings"

	"coop_back/middlewares"
	"coop_back/models"
	"coop_back/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define a struct for the model (e.g., User)

// Initialize the database connection using GORM
func initDatabase(user string, pass string, dbname string) (*gorm.DB, error) {
	// Define the MySQL connection string
	dataconnect := []string{user, ":", pass, "@tcp(127.0.0.1:3306)/", dbname, "?charset=utf8mb4&parseTime=True&loc=Local"}
	dsn := strings.Join(dataconnect, "")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // DisableForeignKeyConstraintWhenMigrating: true
	if err != nil {
		return nil, err
	}
	// DisableForeignKeyConstraintWhenMigrating: false
	db.AutoMigrate(&models.Faculty{})
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Major{})
	db.AutoMigrate(&models.Training{}, &models.Incharge{}, &models.Job{}, &models.User{}, &models.Role{}, &models.Major{},
		&models.Semester{}, &models.Mooban{}, &models.Tambon{}, &models.Incharge{})
	db.AutoMigrate(&models.Record{})
	db.AutoMigrate(&models.Entrepreneur{}, &models.Mooban{})
	db.AutoMigrate(&models.Plan{})

	// db.AutoMigrate(&models.User{}, &models.Role{})

	return db, nil
}

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		// Trust the proxy headers (X-Forwarded-For, X-Real-IP)
		EnableTrustedProxyCheck: false,
		// TrustedProxies:          []string{},
	})

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// log.Println("Client IP " + os.Getenv("CLIENT_IP"))

	app.Use(cors.New(cors.Config{
		// AllowOriginsFunc: func(ctx *fiber.Ctx) bool {
		// 	return ctx.Origin() == "http://localhost:6007" || ctx.Origin() == "https://coop.ict.up.ac.th"
		// },
		AllowOrigins:     "https://coop.ict.up.ac.th, http://localhost:6007,http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept", // Update with necessary headers
		AllowMethods:     "GET, POST",                    // Update with your allowed methods
		AllowCredentials: true,
	}))

	app.Use(middlewares.CustomDomainMiddleware)

	user := os.Getenv("user")
	pass := os.Getenv("pass")
	dbname := os.Getenv("dbname")

	// Initialize the database
	db, err := initDatabase(user, pass, dbname)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Middleware to pass the database connection to routes
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Register the routes
	routers.UserRoutes(app)
	routers.FacultyRoutes(app)
	routers.MajorRoutes(app)
	routers.RecordRoutes(app)
	routers.TrainingRoutes(app)
	routers.AmphurRoutes(app)
	routers.EntrepreneurRoutes(app)
	routers.FacultyRoutes(app)
	routers.InchargeRoutes(app)
	routers.JobRoutes(app)
	routers.MoobanRoutes(app)
	routers.ProvinceRoutes(app)
	routers.SemesterRoutes(app)
	routers.TambonRoutes(app)
	routers.RoleRoutes(app)
	routers.PlanRoutes(app)
	// Define routes
	// app.Get("/faculty", getFaculty)          // SELECT all users

	// Start the Fiber app
	log.Fatal(app.Listen(":6008"))
}
