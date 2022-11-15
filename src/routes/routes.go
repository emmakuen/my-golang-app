package routes

import (
	"ambassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/admin/register", controllers.Register)
}
