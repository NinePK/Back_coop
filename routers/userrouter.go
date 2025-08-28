package routers

import (
	"coop_back/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user_route := app.Group("/user")

	user_route.Get("/", controllers.GetUsers) // SELECT all users
	user_route.Get("/search", controllers.GetUsersOption)
	user_route.Get("/:id", controllers.GetUser)        // SELECT user by ID
	user_route.Post("/", controllers.CreateUser)       // INSERT new user
	user_route.Post("/update", controllers.UpdateUser) // UPDATE user by ID
	user_route.Post("/delete", controllers.DeleteUser) // DELETE user by ID

	user_route.Post("/search", controllers.SearchUser) // INSERT new user
	user_route.Get("/search/:username", controllers.SearchUserByUsername)
}
