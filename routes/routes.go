package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/blogbackend/controller"
	"github.com/kingztech2019/blogbackend/middleware"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate)
	app.Post("api/post", controller.CreatePost)
}
