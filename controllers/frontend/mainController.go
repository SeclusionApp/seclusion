package frontend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seclusionapp/seclusion/models"
	"github.com/seclusionapp/seclusion/views"
)

func Index(c *fiber.Ctx) error {

	//Model
	user := models.GetUser(c) // Add more data as needed
	data := []interface{}{user}

	//View
	return views.Index(c, data)

}
