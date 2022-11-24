package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/seclusionapp/seclusion/config"
	"github.com/seclusionapp/seclusion/controllers/api/errors"
	db "github.com/seclusionapp/seclusion/database"
	"github.com/seclusionapp/seclusion/routes"
)

func main() {
	db.Connect()
	app := fiber.New()
	
	app.Use(
		requestid.New(),
		limiter.New(limiter.Config{
			Max:          config.MAX_REQUESTS,
			LimitReached: errors.LimitReached,
		}),

		logger.New(*config.LOGGER),
		cors.New(*config.CORS),
		
	)
	routes.Setup(app)

	log.Println("[INFO] Server started on port " + config.PORT)
	log.Println("[INFO] Accessing database: " + config.DB_NAME)
	
	log.Fatal(app.Listen(config.PORT))
}
