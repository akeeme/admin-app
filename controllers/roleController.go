package controllers

import (
	"strconv"
	"github.com/akeeme/admin-app/database"
	"github.com/akeeme/admin-app/models"
	"github.com/gofiber/fiber/v2"
)


// CRUD = Create, Read, Update, Delete methods for Users

func AllRoles(c *fiber.Ctx) error {
	// get all Roles

	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}


func CreateRole(c *fiber.Ctx) error {
	// create a Role
	var roles models.Role

	if err := c.BodyParser(&roles); err != nil {
		return err
	}

	database.DB.Create(&roles)

	return c.JSON(roles)

}

func GetRole(c *fiber.Ctx) error {
	// get a role

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	// update a role

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}


func DeleteRole(c *fiber.Ctx) error {
	// delete a user

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return c.JSON(fiber.Map{
		"message": "Role deleted",
	})
}