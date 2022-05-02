package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/seclusionapp/seclusion/config"
	db "github.com/seclusionapp/seclusion/database"
	"github.com/seclusionapp/seclusion/routes"
	"github.com/seclusionapp/seclusion/util"
)

func main() {
	db.Connect()
	PORT := util.GetPort()
	app := fiber.New()
	routes.Setup(app)

	file, err := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	util.HandleError(err, "Failed to open log file")
	defer file.Close()

	app.Use(
		requestid.New(),
		limiter.New(limiter.Config{
			Max: 5,
			LimitReached: func(c *fiber.Ctx) error {
				return c.SendFile("./pages/toofast.html")
			},
		}),
		logger.New(logger.Config{
			Format:     config.LOGGER_FORMAT,
			TimeFormat: config.LOGGER_TIME_FORMAT,
			TimeZone:   config.LOGGER_TIME_ZONE,
			Output:     file,
		}),
		cors.New(cors.Config{
			AllowCredentials: true,
		}),
	)
	log.Fatal(app.Listen(PORT))
}
