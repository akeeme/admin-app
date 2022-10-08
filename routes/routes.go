package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/akeeme/admin-app/controllers"
	"github.com/akeeme/admin-app/middlewares"
)

func Setup(app *fiber.App) {
	
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated) // so we can only access the routes below if we are authenticated
	
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)


	// all CRUD routes for users
	app.Get("/api/users", controllers.AllUsers) // get all users
	app.Post("/api/users", controllers.CreateUser) // create a user
	app.Get("/api/users/:id", controllers.GetUser) // get a user by id
	app.Put("/api/users/:id", controllers.UpdateUser) // update a user by id
	app.Delete("/api/users/:id", controllers.DeleteUser) // delete a user by id

	// all CRUD routes for roles
	app.Get("/api/roles", controllers.AllRoles) // get all roles
	app.Post("/api/roles", controllers.CreateRole) // create a Role
	app.Get("/api/roles/:id", controllers.GetRole) // get a Role by id
	app.Put("/api/roles/:id", controllers.UpdateRole) // update a Role by id
	app.Delete("/api/roles/:id", controllers.DeleteRole) // delete a Role by id
}