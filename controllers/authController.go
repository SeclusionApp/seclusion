package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/seclusionapp/seclusion/config"
	"github.com/seclusionapp/seclusion/database"
	"github.com/seclusionapp/seclusion/models"
	"github.com/seclusionapp/seclusion/util"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if username is taken
	user := models.User{}

	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Username is already taken",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	util.HandleError(err, "Failed to hash password")

	user = models.User{
		Username: data["username"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := models.User{}

	database.DB.Where("username = ?", data["username"]).First(&user)

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Second * config.JWT_EXPIRY).Unix(),
	})

	token, err := claims.SignedString([]byte(util.GetEnv("JWT_SECRET", "secret")))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to sign token",
		})
	}

	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Second * config.JWT_EXPIRY),
		HTTPOnly: true,
	}

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
		"token":   token,
	})
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
