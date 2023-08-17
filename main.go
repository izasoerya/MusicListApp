package main

import (
	"fmt"
	"music_list/database"
	"music_list/models"
	"music_list/route"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	route.MusicRoute(app.Group("/music"))
	route.AuthRoute(app.Group("/auth"))
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":2000"
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
   	 	panic("Error loading .env file")
  	}
}

func main() {
	// loadEnv()
	err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err.Error())
	}

	if err := database.DB.AutoMigrate(&models.Music{}, &models.Auth{}); err != nil {
		fmt.Println("Error automigrate table")
		panic(err.Error())
	}

	app := fiber.New()
	setupRoutes(app)

	app.Listen(getPort())
}
