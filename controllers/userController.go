package controllers

import (
	"strconv"
	"github.com/akeeme/admin-app/database"
	"github.com/akeeme/admin-app/models"
	"github.com/gofiber/fiber/v2"
)


// CRUD = Create, Read, Update, Delete methods for Users

func AllUsers(c *fiber.Ctx) error {
	// get all users

	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}


func CreateUser(c *fiber.Ctx) error {
	// create a user
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)

}

func GetUser(c *fiber.Ctx) error {
	// get a user

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	// update a user

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}


func DeleteUser(c *fiber.Ctx) error {
	// delete a user

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return c.JSON(fiber.Map{
		"message": "User deleted",
	})
}