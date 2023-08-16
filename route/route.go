package route

import (
	"music_list/controller"

	"github.com/gofiber/fiber/v2"
)


func MusicRoute(route fiber.Router) {
	route.Get("", controller.ShowList)
	route.Get("/:id", controller.SearchList)
	route.Post("", controller.AppendList)
	route.Put("/:id", controller.EditList)
	route.Delete("/:id", controller.DeleteList)
}

func AuthRoute(route fiber.Router) {
	route.Get("", controller.CheckUser)
	route.Post("", controller.CreateUser)
	route.Put("/:id", controller.ChangeUser)
}