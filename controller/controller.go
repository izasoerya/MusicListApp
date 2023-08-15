package controller

import (
	"fmt"
	"music_list/database"
	"music_list/models"

	"github.com/gofiber/fiber/v2"
)

func ShowList(c *fiber.Ctx) error {
	var musicList []models.Music
	if err := database.DB.Find(&musicList).Error; err != nil {
		fmt.Println("Error querying music data")
		panic(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success send all data",
		"data": fiber.Map{
			"list": musicList,
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
	list := new(models.Music)
	list.Title = &body.Title
	list.Author = &body.Author
	list.Date = &body.Date
	list.Link = &body.Link

	// Push to Database
	database.DB.Create(&list)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Data created successfully",
		"data": fiber.Map{
			"list": list,
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
