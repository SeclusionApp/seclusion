package errors

import "github.com/gofiber/fiber/v2"

func LimitReached(c *fiber.Ctx) error {
	return c.SendFile("/mnt/c/Users/madga/OneDrive/Desktop/SimpleUserLogin/seclusion/controllers/api/errors/toofast.txt")
}
