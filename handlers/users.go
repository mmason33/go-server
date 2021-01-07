package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmason33/go-project/database"
)

func UsersPage(c *fiber.Ctx) error {
	return c.Render("users", fiber.Map{})
}

// UserGet returns a user
func UserList(c *fiber.Ctx) error {
	users := database.Get("users")
	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	user := map[string]interface{}{
		"name": c.FormValue("name"),
		"age":  c.FormValue("age"),
	}

	database.Insert("users", user)
	return c.JSON(fiber.Map{
		"success": true,
		"user":    "user",
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
