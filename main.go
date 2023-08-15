package main

import (
	"fmt"
	"music_list/database"
	"music_list/route"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	route.MusicRoute(app.Group("/music"))
}

func main() {
	err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err.Error())
	}

	app := fiber.New()
	setupRoutes(app)

	app.Listen(":2000")
}
