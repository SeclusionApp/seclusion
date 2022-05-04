package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	"github.com/seclusionapp/seclusion/config"
	db "github.com/seclusionapp/seclusion/database"
	"github.com/seclusionapp/seclusion/routes"
	"github.com/seclusionapp/seclusion/util"
)

func main() {
	db.Connect()
	PORT := util.GetPort()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	routes.Setup(app)

	app.Use(
		requestid.New(),
		limiter.New(limiter.Config{
			Max: 5,
			LimitReached: func(c *fiber.Ctx) error {
				return c.SendFile("./public/toofast.html")
			},
		}),
		logger.New(*config.LOGGER),
		cors.New(cors.Config{
			AllowCredentials: true,
		}),
	)
	log.Fatal(app.Listen(PORT))
}
