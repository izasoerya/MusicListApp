package controller

import (
	"fmt"
	"music_list/models"

	"github.com/gofiber/fiber/v2"
)
func HomePage(c *fiber.Ctx) error {
	fmt.Println("Hello world!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
	})
}

func ShowList(c *fiber.Ctx) error {
	list := &models.Music{}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
		"data": fiber.Map{
			"list" : list,
		},
	})
}

func AppendList(c *fiber.Ctx) error {
	fmt.Println("Hello world!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
	})
}

func EditList(c *fiber.Ctx) error {
	fmt.Println("Hello world!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
	})
}

func DeleteList(c *fiber.Ctx) error {
	fmt.Println("Hello world!")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
	})
}