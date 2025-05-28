package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type RouterData struct {
	RouterName     string
	ControllerName string
}

var pluralRules = []struct {
	pattern     string
	replacement string
}{
	{pattern: "y$", replacement: "ies"},   // ends with y, replace with ies
	{pattern: "s$", replacement: "es"},    // ends with s, replace with es
	{pattern: "ch$", replacement: "ches"}, // ends with ch, replace with ches
	{pattern: "sh$", replacement: "shes"}, // ends with sh, replace with shes
	// Add more rules as needed
}

func pluralize(word string) string {
	for _, rule := range pluralRules {
		if strings.HasSuffix(word, rule.pattern) {
			return strings.Replace(word, rule.pattern, rule.replacement, 1)
		}
	}
	return word + "s"
}

func main() {
	// List of router names
	routerNames := []string{"plan"} // "role" ,"amphur", "entrepreneur", "faculty", "incharge", "job", "mooban", "province", "semester", "tambon"

	// Loop through each name and generate the router file
	for _, router := range routerNames {
		data := RouterData{
			RouterName:     router,            // "Faculty", "Student", etc.
			ControllerName: pluralize(router), // "Faculties", "Students", etc. (assuming plural controllers)
		}

		// Call the function to generate the file
		err := generateRouterFile(data)
		if err != nil {
			fmt.Println("Error generating router file for", router, ":", err)
		} else {
			fmt.Println("Generated router file for", router)
		}

		err = generateControllerFile(data)
		if err != nil {
			fmt.Println("Error generating controller file for", router, ":", err)
		} else {
			fmt.Println("Generated controller file for", router)
		}
	}
}

// Function to generate a router file
func generateRouterFile(data RouterData) error {
	// Define the template for the router file
	const routerTemplate = `package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func {{upper .RouterName}}Routes(app *fiber.App) {
	{{lower .RouterName}}_route := app.Group("/{{lower .RouterName}}")

	{{lower .RouterName}}_route.Get("/", controllers.Get{{upper .ControllerName}})      // SELECT all {{.RouterName}}s
	{{lower .RouterName}}_route.Get("/:id", controllers.Get{{upper .RouterName}})       // SELECT {{.RouterName}} by ID
	{{lower .RouterName}}_route.Post("/", controllers.Create{{upper .RouterName}})      // INSERT new {{.RouterName}}
	{{lower .RouterName}}_route.Post("/:id", controllers.Update{{upper .RouterName}})   // UPDATE {{.RouterName}} by ID
	{{lower .RouterName}}_route.Post("/:id", controllers.Delete{{upper .RouterName}})   // DELETE {{.RouterName}} by ID
}`

	// Create a template instance and parse the template
	tmpl, err := template.New("router").Funcs(template.FuncMap{
		"lower": func(s string) string { return strings.ToLower(string(s[0]|' ')) + s[1:] },
		"upper": func(s string) string { return strings.ToUpper(string(s[0]|' ')) + s[1:] },
	}).Parse(routerTemplate)
	if err != nil {
		return err
	}

	// Create the file under the routers directory
	fileName := fmt.Sprintf("routers/%srouter.go", data.RouterName)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute the template and write the output to the file
	return tmpl.Execute(file, data)
}

func generateControllerFile(data RouterData) error {
	// Define the template for the router file
	const controllerTemplate = `package controllers

	import (
		"coop_back/models"
	
		"github.com/gofiber/fiber/v2"
	
		"gorm.io/gorm"
	)
	
	// Get all {{upper .RouterName}} (SELECT)
	func Get{{upper .ControllerName}}(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		var {{lower .ControllerName}} []models.{{upper .RouterName}}
		if result := db.Find(&{{lower .ControllerName}}); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON({{lower .ControllerName}})
	}
	
	// Get {{upper .RouterName}} by ID (SELECT)
	func Get{{upper .RouterName}}(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
		var {{upper .RouterName}} models.{{upper .RouterName}}
		if result := db.First(&{{upper .RouterName}}, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "{{upper .RouterName}} not found"})
		}
		return c.JSON({{upper .RouterName}})
	}
	
	// Create new {{upper .RouterName}} (INSERT)
	func Create{{upper .RouterName}}(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		{{upper .RouterName}} := new(models.{{upper .RouterName}})
		if err := c.BodyParser({{upper .RouterName}}); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		if result := db.Create(&{{upper .RouterName}}); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
		return c.JSON({{upper .RouterName}})
	}
	
	// Update existing {{upper .RouterName}} (UPDATE)
	func Update{{upper .RouterName}}(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var {{upper .RouterName}} models.{{upper .RouterName}}
		if result := db.First(&{{upper .RouterName}}, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "{{upper .RouterName}} not found"})
		}
	
		// Parse request body into struct
		if err := c.BodyParser(&{{upper .RouterName}}); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	
		db.Save(&{{upper .RouterName}})
		return c.JSON({{upper .RouterName}})
	}
	
	// Delete {{upper .RouterName}} by ID (DELETE)
	func Delete{{upper .RouterName}}(c *fiber.Ctx) error {
		id := c.Params("id")
		db := c.Locals("db").(*gorm.DB)
	
		var {{upper .RouterName}} models.{{upper .RouterName}}
		if result := db.First(&{{upper .RouterName}}, id); result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "{{upper .RouterName}} not found"})
		}
	
		db.Delete(&{{upper .RouterName}})
		return c.JSON(fiber.Map{"message": "{{upper .RouterName}} deleted successfully"})
	}
	`

	// Create a template instance and parse the template
	tmpl, err := template.New("router").Funcs(template.FuncMap{
		"lower": func(s string) string { return strings.ToLower(string(s[0]|' ')) + s[1:] },
		"upper": func(s string) string { return strings.ToUpper(string(s[0]|' ')) + s[1:] },
	}).Parse(controllerTemplate)
	if err != nil {
		return err
	}

	// Create the file under the routers directory
	fileName := fmt.Sprintf("controllers/%scontroller.go", data.RouterName)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute the template and write the output to the file
	return tmpl.Execute(file, data)
}
