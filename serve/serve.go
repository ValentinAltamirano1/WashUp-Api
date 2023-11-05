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
	app.Post("/social-login", handler.UserSocialLogin)
	app.Post("/reservations", handler.ReservaCreate)
	app.Get("/fechasdisponibles/:service", handler.ObtenerFechasDisponiblesHandler)
	app.Get("/horariosdisponibles/{service}{time}", handler.ObtenerHorariosDisponiblesHandler)
	app.Delete("/employee/delete", handler.EmployeeDelete)
	app.Get("/horariosdisponibles/:service/:date", handler.ObtenerHorariosDisponiblesHandler)
	app.Post("/crear-preferencia", handler.PaymentMercadoPago)
	app.Get("/employee/reservations-without-assignment", handler.GetAllReservationsWithoutEmployee)
	app.Post("/employee/confirm-reservation", handler.EmployeeConfirmReservation)
	app.Get("/employee/reservations/assigned/:email", handler.GetAllReservationsByEmployee)
	app.Post("/employee/reservation-done", handler.EmployeeReservationDone)
	app.Get("/employee/reservations/done/assigned/:email", handler.GetAllReservationsDoneByEmployee)
	app.Get("/my-reservations/:userID", handler.ObtenerMisReservas)

	
	app.Use(middleware.AuthRequired())

	return app
}
