package handler

import (
	"fmt"

	"github.com/ValentinAltamirano1/WashUp-Api/database"
	"github.com/ValentinAltamirano1/WashUp-Api/model"
	"github.com/ValentinAltamirano1/WashUp-Api/service"
	"github.com/gofiber/fiber/v2"
)

func ObtenerFechasDisponiblesHandler(c *fiber.Ctx) error {
    servicio := c.Params("servicio") // Obtiene el servicio de la solicitud (ajusta esto según tu enrutamiento)
	fmt.Println("servicio")
	fmt.Println(servicio)
    // Obtén el servicio de reserva
    db := database.DB
    reservationClient := model.ReservationClient{DB: db}

    // Llama a la función service.ObtenerFechasDisponibles para obtener las fechas no disponibles.
    fechasNoDisponibles, err := service.ObtenerFechasDisponibles(reservationClient, servicio)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "error al obtener fechas no disponibles",
        })
    }

    // Responde con un código 200 (OK) y envía las fechas no disponibles.
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "fechas_no_disponibles": fechasNoDisponibles,
    })
}


// ReservaCreate maneja las solicitudes para crear una nueva reserva.
func ReservaCreate(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	var params service.ReservationParams

	// Analiza los datos JSON enviados en la solicitud HTTP y los almacena en la variable params.
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al analizar JSON",
		})
	}

	// Llama a la función service.CreateReservation para crear una nueva reserva.
	reserva, err := service.CreateReservation(reservationClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al crear reserva",
		})
	}

	// Responde con un código 201 (Creado) y envía los detalles de la reserva creada.
	return c.Status(fiber.StatusCreated).JSON(reserva)
}

// ReservaCheck maneja las solicitudes para verificar la disponibilidad de un horario y fecha específicos.
func ReservaCheck(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	var params service.ReservationCheckParams

	// Analiza los datos JSON enviados en la solicitud HTTP y los almacena en la variable params.
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al analizar JSON",
		})
	}

	// Llama a la función service.CheckReservation para verificar la disponibilidad del horario y fecha.
	disponible, err := service.CheckReservation(reservationClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al verificar disponibilidad",
		})
	}

	// Responde con un código 200 (Éxito) y envía un valor booleano que indica si el horario está disponible.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"disponible": disponible,
	})
}
