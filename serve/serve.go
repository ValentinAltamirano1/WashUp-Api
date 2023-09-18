package serve

import (
	"github.com/ValentinAltamirano1/WashUp-Api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetRouters() *fiber.App {
	app := fiber.New()

	app.Get("/users", handler.UserCreate)

	return app
}
