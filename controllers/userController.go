package controllers

// CRUD = Create, Read, Update, Delete

import (
	"github.com/akeeme/admin-app/database"
	"github.com/akeeme/admin-app/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func AllUsers(c *fiber.Ctx) error {
	// get all users

	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}


func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}


	password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14) // so we store password as a hash

	user.Password = password

	database.DB.Create(&user)

	return c.JSON(user)

}