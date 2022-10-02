package main

import (
    "github.com/akeeme/admin-app/database"
    "github.com/akeeme/admin-app/routes"
    "github.com/gofiber/fiber/v2"
    
)



func main() {
    database.Connect()

    app := fiber.New()

    routes.Setup(app)

    app.Listen(":8000")


   
}