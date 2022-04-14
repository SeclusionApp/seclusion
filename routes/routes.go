package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/controllers"
)

func Setup(app *fiber.App) {

	/*
	 * These are the routes for the main API.
	 * Note that the routes are protected by JWT.
	 */

	/*
	 * Main Routes
	 */
	app.Get("/v1/", controllers.Index)

	/*
	 * Auth Routes
	 */
	app.Post("/v1/auth/register", controllers.Register) // Register a new user
	app.Post("/v1/auth/login", controllers.Login)       // Login a user
	app.Post("/v1/auth/logout", controllers.Logout)     // Logout a user

	/*
	 * User Routes
	 */
	app.Get("/v1/user", controllers.User)                                // Get User
	app.Get("/v1/user/channels", controllers.UserChannels)               // Get User Channels
	app.Get("/v1/user/channels/:channel_id", controllers.UserChannel)    // Get User Channel by ID
	app.Post("/v1/user/channels/:channel_id", controllers.UserChannel)   // Add User to Channel
	app.Delete("/v1/user/channels/:channel_id", controllers.UserChannel) // Remove User from Channel

}
