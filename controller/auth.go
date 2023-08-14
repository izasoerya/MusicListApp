package controller

import (
	"music_list/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	account := &models.Auth{}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
		"data": fiber.Map{
			"account" : account,
		},
	})
}