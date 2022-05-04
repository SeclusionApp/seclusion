package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/database"
	models "github.com/seclusionapp/seclusion/database/structs"
)

func GetMessages(c *fiber.Ctx) error {
	var messages []models.Message

	if c.Params("channel_id") != "" {
		database.DB.Where("channel_id = ?", c.Params("channel_id")).Find(&messages) // Get messages from database
	} else {
		database.DB.Find(&messages) // Get messages from database
	}
	return c.JSON(fiber.Map{
		"status":   "ok",
		"method":   c.Method(),
		"messages": messages,
	})
}
