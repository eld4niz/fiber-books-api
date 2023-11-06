package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/eld4niz/fiber-books-api/database"
    "github.com/eld4niz/fiber-books-api/router"
)

func main() {
    // Start a new fiber app
    app := fiber.New()

    // Connect to the Database
    database.ConnectDB()

    // Setup the routes
    router.SetupRoutes(app)

    // Listen on PORT 3000
    app.Listen(":3000")
}
