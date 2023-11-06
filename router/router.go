package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	bookRoutes "github.com/eld4niz/fiber-books-api/internal/routes/books"
)
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	bookRoutes.SetupBookRoutes(api)
	
}
