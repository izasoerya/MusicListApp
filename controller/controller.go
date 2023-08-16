package controller

import (
	"fmt"
	"music_list/database"
	"music_list/models"
	"strconv"

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
	var body models.Request
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	// Append models to data
	list := new(models.Music)
	list.ID = &body.ID
	list.Title = &body.Title
	list.Author = &body.Author
	list.Date = &body.Date
	list.Link = &body.Link

	// Push to Database
	database.DB.Create(&list)

	fmt.Println("Successfully created")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Data created successfully",
		"data": fiber.Map{
			"list": list,
		},
	})
}

func EditList(c *fiber.Ctx) error {
	// Handle input
	paramID := c.Params("id") // id handler
	ID, err := strconv.Atoi(paramID)
	if err != nil {
		panic(err)
	}
	var body models.Request // body handler
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	// Edit the DB
	var data models.Music
	if err := database.DB.Model(&data).Where("id = ?", ID).Updates(models.Music{
		ID : &body.ID,
		Author: &body.Author,
		Title:  &body.Title,
		Date: &body.Date,
	}).Error; err != nil {
		fmt.Println("Error updating music data")
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

func DeleteList(c *fiber.Ctx) error {
	paramID := c.Params("id") // id handler
	ID, err := strconv.Atoi(paramID)
	if err != nil {
		panic(err)
	}
	fmt.Println(ID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
	})
}
