package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/models"
)

func UserChannel(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	channel := c.Locals("channel").(*models.Channel)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "User Channel Hit",
		"user":    user,
		"channel": channel,
	})
}

func UserChannels(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "User Channels Hit",
		"user":    user,
	})
}
