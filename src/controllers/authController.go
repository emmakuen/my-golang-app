package controllers

import "github.com/gofiber/fiber/v2"

func Register(context *fiber.Ctx) error {
	return context.JSON(fiber.Map{
		"message": "Hello",
	})
}
