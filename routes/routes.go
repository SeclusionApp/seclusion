package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/seclusionapp/seclusion/controllers/api"
	"github.com/seclusionapp/seclusion/controllers/frontend"
)

func Setup(app *fiber.App) {

	/*
	 * Web Routes
	 *
	 */
	app.Get("/", frontend.Index)

	/*
	 * These are the routes for the main API.
	 * Note that the routes are protected by a JWT token stored in a cookie after auth.
	 */

	/*
	 * Main Routes API
	 */
	app.Get("/v1/", api.Index)
	app.Get("/v1/status", monitor.New(
		monitor.Config{
			APIOnly: true,
			Next:    nil,
		},
	))
	/*
	 * Auth Routes
	 */
	app.Post("/v1/auth/register", api.Register) // Register a new user
	app.Post("/v1/auth/login", api.Login)       // Login a user
	app.Post("/v1/auth/logout", api.Logout)     // Logout a user

	/*
	 * User Routes
	 */
	app.Get("/v1/user", api.User)
	app.Get("/v1/users/:id", api.User)

	/*
	 * Channel Routes
	 */
	// Getters
	app.Get("/v1/channels", api.Channels)
	app.Get("/v1/channels/:id", api.Channel)

	// Creators
	app.Post("/v1/channels/:name", api.Channel)

	// Deletions
	app.Delete("/v1/channels/:id", api.Channel)

	/*
	 * Message Routes
	 */
	app.Get("/v1/messages", api.Messages)             // Get all messages
	app.Get("/v1/message/:id", api.Message)           // Get a message by ID
	app.Get("/v1/messages/:channel_id", api.Messages) // Get all messages for a channel
	app.Post("/v1/message", api.Message)              // Create a new message
	app.Delete("/v1/message/:id", api.Message)        // Delete a message by ID

}
