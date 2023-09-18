package handler

import (


	"github.com/ValentinAltamirano1/WashUp-Api/database"
	"github.com/ValentinAltamirano1/WashUp-Api/model"
	"github.com/ValentinAltamirano1/WashUp-Api/service"
	"github.com/gofiber/fiber/v2"
)

func UserCreate(c *fiber.Ctx) error {

	db := database.DB
	userClient := model.UserClient{DB: db}
	params := service.UserParams{
		Name: c.Query("name"),
		Email: c.Query("email"),
		Password: c.Query("password"),
	}

	err := service.CreateUser(userClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to create user",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}