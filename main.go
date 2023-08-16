package main

import (
	"fmt"
	"log"
	"music_list/database"
	"music_list/models"
	"music_list/route"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	route.MusicRoute(app.Group("/music"))
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err.Error())
	}
	database.DB.AutoMigrate(&models.Music{})

	app := fiber.New()
	setupRoutes(app)

	app.Listen(":2000")
}
