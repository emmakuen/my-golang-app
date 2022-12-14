package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(context *fiber.Ctx) error {
	var data map[string]string

	if err := context.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		context.Status(400)
		return context.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		Firstname:    data["firstname"],
		Lastname:     data["lastname"],
		Email:        data["email"],
		Password:     password,
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return context.JSON(user)
}

func Login(context *fiber.Ctx) error {
	var data map[string]string

	if err := context.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		context.Status(fiber.StatusBadRequest)

		return context.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		context.Status(fiber.StatusBadRequest)

		return context.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	return context.JSON(user)
}
