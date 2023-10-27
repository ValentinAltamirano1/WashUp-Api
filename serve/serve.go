package serve

import (
	"github.com/ValentinAltamirano1/WashUp-Api/handler"
	"github.com/ValentinAltamirano1/WashUp-Api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetRouters() *fiber.App {
	app := fiber.New()

	// Configurar encabezados CORS para permitir solicitudes desde http://localhost:3000
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept,Authorization",
		AllowCredentials: true,
	}))

	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Post("/users", handler.UserCreate)
	app.Post("/login", handler.UserLogin)
	app.Post("/employee", handler.EmployeeCreate)
	app.Post("/employeelogin", handler.EmployeeLogin)

	app.Use(middleware.AuthRequired())

	return app
}
