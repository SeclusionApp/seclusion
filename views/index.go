package views

import (
	"github.com/gofiber/fiber/v2"
	structs "github.com/seclusionapp/seclusion/database/structs"
)

func Index(c *fiber.Ctx, data []interface{}) error {
	//Find user in data
	user := data[0].(*structs.User) // TODO: test for user info leak
	if user.ID == 0 {
		user.Username = "Guest"
	}
	return c.Render("templates/index", fiber.Map{
		"Username": user.Username,
	})
}
