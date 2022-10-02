package database

import (
	"github.com/joho/godotenv"
    "os"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
	"github.com/akeeme/admin-app/models"
    _ "fmt"
)

var DB *gorm.DB

func getENV(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

func Connect() {
	db, err := gorm.Open(mysql.Open(getENV("dsn")), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

	DB = db

	db.AutoMigrate(&models.User{})
}