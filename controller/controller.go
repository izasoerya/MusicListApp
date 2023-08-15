package controller

import (
	"fmt"
	"music_list/database"
	"music_list/models"

	"github.com/gofiber/fiber/v2"
)

func ShowList(c *fiber.Ctx) error {
	var data []models.Music = make([]models.Music, 0)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
		"data": fiber.Map{
			"list": data,
		},
	})
}

func AppendList(c *fiber.Ctx) error {
	// Parse body input
	type Request struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Date   string `json:"date"`
		Link   string `json:"link"`
	}
	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
	}

	// Append models to data
	data := new(models.Music)
	data.Author = &body.Author
	data.Title = &body.Title
	data.Date = &body.Date
	data.Link = &body.Link

	// Push to Database
	result := database.DB.Create(&data)
	if result != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Data created successfully",
		"data": fiber.Map{
			"data": data,
		},
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
