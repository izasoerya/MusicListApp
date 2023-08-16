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
	route.AuthRoute(app.Group("/auth"))
}

func getPort() string {
	if envPort := "MY_PORT"; envPort != "" {
		return ":" + envPort
	}
	return ":2000"
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

	if err := database.DB.AutoMigrate(&models.Music{}); err != nil {
		fmt.Println("Error creating 'music' table")
		panic(err.Error())
	}
	if err := database.DB.AutoMigrate(&models.Auth{}); err != nil {
		fmt.Println("Error creating 'auth' table")
		panic(err.Error())
	}

	app := fiber.New()
	setupRoutes(app)

	app.Listen(getPort())
}
