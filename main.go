package main

import (
	"music_list/route"
	"music_list/database"
	"music_list/models"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	route.MusicRoute(app.Group("/music"))
}

func main() {
	db, err := database.ConnectDB()
	if err != nil {
        panic(err.Error())
    }
	db.AutoMigrate(&models.Music{})

	app := fiber.New()
	setupRoutes(app)
	
	app.Listen(":2000")
}
