package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db "github.com/seclusionapp/seclusion/database"
	"github.com/seclusionapp/seclusion/routes"
	"github.com/seclusionapp/seclusion/util"
)

func main() {
	db.Connect()
	util.InitEnv()

	PORT := util.GetPort()

	app := fiber.New()

	routes.Setup(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	log.Fatal(app.Listen(PORT))
}
