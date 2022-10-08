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
}