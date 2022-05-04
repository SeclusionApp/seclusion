package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/seclusionapp/seclusion/database"
	models "github.com/seclusionapp/seclusion/database/structs"
	"github.com/seclusionapp/seclusion/util"
)

func GetUser(c *fiber.Ctx) *models.User {
	token := c.Cookies("token")
	user := &models.User{}
	if !util.VerifyToken(token) {
		return user // Not logged in
	}
	tokenParse := util.GetToken(token)
	claims := tokenParse.Claims.(*jwt.StandardClaims)
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	return user
}
