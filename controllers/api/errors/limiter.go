package errors

import "github.com/gofiber/fiber/v2"

func LimitReached(c *fiber.Ctx) error {
	return c.SendFile("./public/toofast.html")
}
