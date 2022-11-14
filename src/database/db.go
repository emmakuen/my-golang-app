package database

import (
	"ambassador/src/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	godotenv.Load()

	DB, err = gorm.Open(mysql.Open("root:"+os.Getenv("DB_ROOT_PASSWORD")+"@tcp(db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database!")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
