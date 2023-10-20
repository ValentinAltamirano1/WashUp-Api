package serve

import (
	"github.com/ValentinAltamirano1/WashUp-Api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetRouters() *fiber.App {
	app := fiber.New()

	// Configurar encabezados CORS para permitir solicitudes desde http://localhost:3000
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Set("Access-Control-Allow-Credentials", "true")
		return c.Next()
	})

	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Post("/users", handler.UserCreate)
	app.Post("/login", handler.UserLogin)

	return app
}
