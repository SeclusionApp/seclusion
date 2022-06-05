package api

import (
	"github.com/gofiber/fiber/v2"
)

func Reroute(c *fiber.Ctx) error {
	return c.Redirect("/v1/")
}

func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Seclusion Auth API Hit. Welcome!",
	})
}
