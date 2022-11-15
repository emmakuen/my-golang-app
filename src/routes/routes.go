package routes

import (
	"ambassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	baseApi := app.Group("api")

	adminApi := baseApi.Group("admin")

	adminApi.Post("/register", controllers.Register)
	adminApi.Post("/login", controllers.Login)
}
