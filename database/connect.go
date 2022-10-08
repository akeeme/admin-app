package database

import (
	_ "fmt"
	"os"

	"github.com/akeeme/admin-app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db.AutoMigrate(&models.User{}, &models.Role{})
}