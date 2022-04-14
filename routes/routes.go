package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/controllers"
)

func Setup(app *fiber.App) {

	/*
	 * Main Routes
	 */
	app.Get("/v1/", controllers.Index)

	/*
	 * Auth Routes
	 */
	app.Post("/v1/auth/register", controllers.Register)
	app.Post("/v1/auth/login", controllers.Login)
	app.Post("/v1/auth/logout", controllers.Logout)

	/*
	 * User Routes
	 */
	app.Get("/v1/user", controllers.User)
	//app.Get("/v1/:user/channels", controllers.UserChannels)
	//app.Get("/v1/:user/channels/:channel", controllers.UserChannel)

}
