package api

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/database"
	structs "github.com/seclusionapp/seclusion/database/structs"
	"github.com/seclusionapp/seclusion/models"
	"github.com/seclusionapp/seclusion/util"
)

func Messages(c *fiber.Ctx) error {

	token := c.Cookies("token")

	if !util.VerifyToken(token) {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Token",
		})
	}

	if c.Method() == "GET" {
		models.GetMessages(c)
		return nil
	}

	return c.Status(400).JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid request",
	})

}

func Message(c *fiber.Ctx) error {

	token := c.Cookies("token")

	if !util.VerifyToken(token) {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Token",
		})
	}

	if c.Method() == "GET" {

		// if is id param
		if c.Params("id") != "" {
			var message structs.Message
			database.DB.Where("id = ?", c.Params("id")).First(&message) // Get message from database
			return c.JSON(fiber.Map{
				"status":  "ok",
				"method":  c.Method(),
				"message": message,
			})
		}

		// if is channel_id param
		if c.Params("channel_id") != "" {
			var messages []structs.Message
			database.DB.Where("channel_id = ?", c.Params("channel_id")).Find(&messages) // Get messages from database
			return c.JSON(fiber.Map{
				"status":   "ok",
				"method":   c.Method(),
				"messages": messages,
			})
		}

		// if is user_id param
		if c.Params("user_id") != "" {
			var messages []structs.Message
			database.DB.Where("user_id = ?", c.Params("user_id")).Find(&messages) // Get messages from database
			return c.JSON(fiber.Map{
				"status":   "ok",
				"method":   c.Method(),
				"messages": messages,
			})
		}

	}

	if c.Method() == "POST" {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": "error",
			})
		}

		var message structs.Message
		//if content number of chars is over 600
		if len(data["content"]) > 600 {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "Message content is over 600 characters",
			})
		}
		message.Content = data["content"]
		user_id, err := strconv.Atoi(data["user_id"])
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": "error",
				"error":  "Invalid user_id",
				"debug":  err,
				"method": c.Method(),
			})
		}
		message.UserID = user_id
		channel_id, err := strconv.Atoi(data["channel_id"])
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": "error",
				"error":  "Invalid channel_id",
				"debug":  err,
				"method": c.Method(),
			})
		}
		message.ChannelID = channel_id
		message.Time = time.Now().Unix()

		database.DB.Create(&message) // Create message in database

		return c.JSON(fiber.Map{
			"status":  "ok",
			"method":  c.Method(),
			"message": message,
		})

	}

	if c.Method() == "DELETE" {
		var message structs.Message
		database.DB.Where("id = ?", c.Params("id")).First(&message) // Get message from database
		database.DB.Delete(&message)                                // Delete message from database
		return c.JSON(fiber.Map{
			"status":  "ok",
			"method":  c.Method(),
			"message": message,
		})
	}

	return c.Status(400).JSON(fiber.Map{
		"status":  "error",
		"method":  c.Method(),
		"message": "Invalid request / Unsupported method",
	})

}
