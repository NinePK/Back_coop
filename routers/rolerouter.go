package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {
	role_route := app.Group("/role")

	role_route.Get("/", controllers.GetRoles)      // SELECT all roles
	role_route.Get("/:id", controllers.GetRole)       // SELECT role by ID
	role_route.Post("/", controllers.CreateRole)      // INSERT new role
	role_route.Post("/:id", controllers.UpdateRole)   // UPDATE role by ID
	role_route.Post("/:id", controllers.DeleteRole)   // DELETE role by ID
}