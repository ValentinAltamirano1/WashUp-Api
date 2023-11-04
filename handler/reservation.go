package handler

import (
	"fmt"
	"strings"

	"github.com/ValentinAltamirano1/WashUp-Api/database"
	"github.com/ValentinAltamirano1/WashUp-Api/model"
	"github.com/ValentinAltamirano1/WashUp-Api/service"
	"github.com/gofiber/fiber/v2"
)

func ObtenerFechasDisponiblesHandler(c *fiber.Ctx) error {
	fmt.Println("URL completa:", c.OriginalURL())
	servicioParam := c.Params("service")
	servicio := strings.Replace(servicioParam, "%20", " ", -1)

	fmt.Println("servicio:", servicio)

	db := database.DB
	reservationClient := model.ReservationClient{DB: db}

	fechasNoDisponibles, err := service.ObtenerFechasDisponibles(reservationClient, servicio)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al obtener fechas no disponibles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fechas_no_disponibles": fechasNoDisponibles,
	})
}

func ObtenerHorariosDisponiblesHandler(c *fiber.Ctx) error {
	fmt.Println("URL completa:", c.OriginalURL())
	servicioParam := c.Params("service")
	fechaParam := c.Params("date")
	servicio := strings.Replace(servicioParam, "%20", " ", -1)
	
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}

	horariosDisponibles	, err := service.ObtenerHorariosDisponibles(reservationClient, servicio,fechaParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al obtener fechas no disponibles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"horarios": horariosDisponibles,
	})
}

func ObtenerMisReservas(c *fiber.Ctx) error {
	//fmt.Println("URL completa:", c.OriginalURL())
	userIDParam := c.Params("userID")


	db := database.DB
	reservationClient := model.ReservationClient{DB: db}

	misReservas, err := service.ObtenerMisReservas(reservationClient, userIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al obtener las reservas del usuario",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"reservations": misReservas,
	})
}




func ReservaCreate(c *fiber.Ctx) error {
	db := database.DB
	userClient := model.UserClient{DB: db}
	reservationClient := model.ReservationClient{DB: db}
	var params service.ReservationParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al analizar JSON",
		})
	}

	reserva, err := service.CreateReservation(reservationClient, userClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al crear reserva",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(reserva)
}

func ReservaCheck(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	var params service.ReservationCheckParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al analizar JSON",
		})
	}

	disponible, err := service.CheckReservation(reservationClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al verificar disponibilidad",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"disponible": disponible,
	})
}
