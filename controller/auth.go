package controller

import (
	"fmt"
	"music_list/database"
	"music_list/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CheckUser(c *fiber.Ctx) error {
	var account []models.Auth
	if err := database.DB.Find(&account).Error; err != nil {
		fmt.Println("Error querying account data")
		panic(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Account created successfully",
		"data": fiber.Map{
			"account": account,
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	var body models.Auth
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}
	account := new(models.Auth)
	account.Username = body.Username
	account.Password = body.Password

	database.DB.Create(&account)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Account created successfully",
		"data": fiber.Map{
			"account": account,
		},
	})
}

func ChangeUser(c *fiber.Ctx) error {
	// Handle input
	paramID := c.Params("id") // id handler
	id, err := strconv.Atoi(paramID)
	if err != nil {
		panic(err)
	}
	var body models.Auth // body handler
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	// Edit the DB
	var data models.Auth
	if err := database.DB.Model(&data).Where("id = ?", id).Updates(models.Auth{
		Username: body.Username,
		Password: body.Password,
	}).Error; err != nil {
		fmt.Println("Error updating account data")
		panic(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Data edited successfully",
		"data": fiber.Map{
			"data": data,
		},
	})
}
