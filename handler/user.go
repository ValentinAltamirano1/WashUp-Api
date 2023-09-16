package handler

import (
	"github.com/gofiber/fiber/v2"
)

func UserCreate(c *fiber.Ctx) error {
	return c.SendString("OK")
}