package controllers

// CRUD = Create, Read, Update, Delete

import (
	"github.com/akeeme/admin-app/database"
	"github.com/akeeme/admin-app/models"
	"github.com/gofiber/fiber/v2"
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

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)

}