package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	_, err := gorm.Open(mysql.Open("root:"+os.Getenv("DB_ROOT_PASSWORD")+"@tcp(db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	print("root:" + os.Getenv("DB_ROOT_PASSWORD") + "@tcp(db:3306)/ambassador")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	log.Fatal(app.Listen(":3000"))
}
