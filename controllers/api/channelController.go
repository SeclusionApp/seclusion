package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

	tokenParse := util.GetToken(token)
	claims := tokenParse.Claims.(*jwt.StandardClaims)
	var channels []structs.Channel //Get all channels from database "channels" table
	database.DB.Where("user_id = ?", claims.Issuer).Find(&structs.Channel_User{}, &channels)
	return c.JSON(fiber.Map{
		"status":   "ok",
		"channels": channels,
	})

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
		return c.JSON(fiber.Map{
			"status":  "ok",
			"method":  c.Method(),
			"channel": channel,
		})
	}

	if c.Method() == "POST" {
		var channel structs.Channel
		channel.Name = c.Params("name")
		database.DB.Create(&channel) // Create channel in database
		// Add user to channel
		tokenParse := util.GetToken(token)
		claims := tokenParse.Claims.(*jwt.StandardClaims)
		var user structs.User
		database.DB.Where("id = ?", claims.Issuer).First(&user) // Get user from database
		database.DB.Create(&structs.Channel_User{
			ChannelID: channel.ID,
			UserID:    user.ID,
		})
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
