package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/seclusionapp/seclusion/database"
	models "github.com/seclusionapp/seclusion/database/structs"
	"github.com/seclusionapp/seclusion/util"
)

func User(c *fiber.Ctx) error {

	token := c.Cookies("token")

	if !util.VerifyToken(token) { // If token is invalid
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Token",
		})
	}

	tokenParse := util.GetToken(token)
	claims := tokenParse.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user) // Get user from database
	// Get channels from database channel_users table
	var channels []models.Channel
	database.DB.Joins("JOIN channel_users ON channel_users.channel_id = channels.id").Where("channel_users.user_id = ?", user.ID).Find(&channels)
	return c.JSON(fiber.Map{
		"status":   "ok",
		"user":     user,
		"channels": channels,
	})

}
