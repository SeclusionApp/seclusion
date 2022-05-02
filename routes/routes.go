package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/seclusionapp/seclusion/controllers"
)

func Setup(app *fiber.App) {

	/*
	 * These are the routes for the main API.
	 * Note that the routes are protected by a JWT token stored in a cookie after auth.
	 */

	/*
	 * Main Routes
	 */
	app.Get("/v1/", controllers.Index)
	app.Get("/v1/status", monitor.New(
		monitor.Config{
			APIOnly: true,
			Next:    nil,
		},
	))
	/*
	 * Auth Routes
	 */
	app.Post("/v1/auth/register", controllers.Register) // Register a new user
	app.Post("/v1/auth/login", controllers.Login)       // Login a user
	app.Post("/v1/auth/logout", controllers.Logout)     // Logout a user

	/*
	 * User Routes
	 */
	app.Get("/v1/user", controllers.User)
	app.Get("/v1/users/:id", controllers.User)

	/*
	 * Channel Routes
	 */
	// Getters
	app.Get("/v1/channels", controllers.Channels)
	app.Get("/v1/channels/:id", controllers.Channel)

	// Creators
	app.Post("/v1/channels/:name", controllers.Channel)

	// Deletions
	app.Delete("/v1/channels/:id", controllers.Channel)

	/*
	 * Message Routes
	 */
	app.Get("/v1/messages", controllers.Messages)             // Get all messages
	app.Get("/v1/message/:id", controllers.Message)           // Get a message by ID
	app.Get("/v1/messages/:channel_id", controllers.Messages) // Get all messages for a channel
	app.Post("/v1/message", controllers.Message)              // Create a new message
	app.Delete("/v1/message/:id", controllers.Message)        // Delete a message by ID

}
