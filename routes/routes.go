package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
