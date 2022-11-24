package api

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/seclusionapp/seclusion/config"
	"github.com/seclusionapp/seclusion/database"
	structs "github.com/seclusionapp/seclusion/database/structs"
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
	user := structs.User{}

	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Username is already taken",
		})
	}

	//Verify Emial
	if !util.ValidMailAddress(data["email"]) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email is not valid",
		})
	}

	//Check if email is taken
	user = structs.User{}

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email is already in use.",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	util.HandleError(err, "Failed to hash password")

	user = structs.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := structs.User{}

	database.DB.Where("email = ?", data["email"]).First(&user)

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
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

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Second * config.JWT_EXPIRY),
		HTTPOnly: true,
		Path:   "/",
		SameSite: "lax",
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
		"token":   token,
		"expires":time.Now().Add(time.Second * config.JWT_EXPIRY),
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
