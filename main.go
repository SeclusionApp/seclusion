package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	db "github.com/seclusionapp/seclusion/database"
	"github.com/seclusionapp/seclusion/routes"
	"github.com/seclusionapp/seclusion/util"
)

func main() {
	db.Connect()
	util.InitEnv()

	PORT := util.GetPort()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	log.Fatal(app.Listen(PORT))
}
