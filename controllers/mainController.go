package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	fmt.Println("Hello World")

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Seclusion Auth API Hit. Welcome!",
	})
}
