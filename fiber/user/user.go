package user

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) {
	c.SendString("All users")
}
