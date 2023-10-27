package handler

import (
	"github.com/ValentinAltamirano1/WashUp-Api/database"
	"github.com/ValentinAltamirano1/WashUp-Api/model"
	"github.com/ValentinAltamirano1/WashUp-Api/service"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func EmployeeCreate(c *fiber.Ctx) error {
	db := database.DB
	employeeClient := model.EmployeeClient{DB: db}
	var params service.EmployeeParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}
	
	employee, err := service.CreateEmployee(employeeClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to create employee",
		})
	}
	fmt.Println(err)
	return c.Status(fiber.StatusCreated).JSON(employee)
}

func EmployeeLogin(c *fiber.Ctx) error {
	db := database.DB
	employeeClient := model.EmployeeClient{DB: db}
	var params service.EmployeeParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	employee, err := service.LoginEmployee(employeeClient, params.Email, params.Password)
	fmt.Println(employee)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to login employee",
		})
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}