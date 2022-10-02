package main

import (
    "github.com/akeeme/admin-app/database"
    "github.com/akeeme/admin-app/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    
)



func main() {
    database.Connect()

    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowCredentials: true, // so we can send cookies
    })) // so we can make requests from different ports

    routes.Setup(app)

    app.Listen(":8000")


   
}