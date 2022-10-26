package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/database"
	structs "github.com/seclusionapp/seclusion/database/structs"
	"github.com/seclusionapp/seclusion/util"
)

func Channels(c *fiber.Ctx) error {

	token := c.Cookies("token")

	if !util.VerifyToken(token) {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Token",
		})
	}

	//Get all channels from DB and return list

	if c.Method() == "GET" {
		var channels []structs.Channel
		database.DB.Find(&channels)
		return c.Status(200).JSON(channels)
	} else {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Method",
		})
	}
}

func Channel(c *fiber.Ctx) error {

	token := c.Cookies("token")

	if !util.VerifyToken(token) {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Token",
		})
	}

	if c.Method() == "GET" {
		var channel structs.Channel
		database.DB.Where("id = ?", c.Params("id")).First(&channel) // Get channel from database
		if channel.ID == 0 {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Channel not found",
			})
		}
		return c.JSON(fiber.Map{
			"status":  "ok",
			"method":  c.Method(),
			"channel": channel,
		})
	}

	if c.Method() == "POST" {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Check if channel name is taken
		channel := structs.Channel{}

		database.DB.Where("name = ?", data["name"]).First(&channel)
		if channel.ID != 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Channel name is already taken",
			})
		}

		channel = structs.Channel{
			Name: data["name"],
		}

		database.DB.Create(&channel)

		return c.JSON(fiber.Map{
			"status":  "ok",
			"method":  c.Method(),
			"channel": channel,
		})
	}

	if c.Method() == "DELETE" {
		var channel structs.Channel
		database.DB.Where("id = ?", c.Params("id")).First(&channel) // Get channel from database
		database.DB.Delete(&channel)                                // Delete channel from database
		return c.JSON(fiber.Map{
			"status":  "ok",
			"method":  c.Method(),
			"channel": channel,
		})
	}

	return c.Status(400).JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid Method",
	})

}

func ChannelMessages(c *fiber.Ctx) error {
	return c.Status(200).JSON([]string{"channel", "messages"})
}
