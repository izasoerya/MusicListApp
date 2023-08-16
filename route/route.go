package route

import (
	"music_list/controller"

	"github.com/gofiber/fiber/v2"
)


func MusicRoute(route fiber.Router) {
	route.Get("/show", controller.ShowList)
	route.Post("/add", controller.AppendList)
	route.Put("/edit/:id", controller.EditList)
	route.Delete("/delete", controller.DeleteList)
}

func AuthRoute(route fiber.Router) {
	route.Post("/create", controller.CreateUser)
}