package bookRoutes

import (
	"github.com/gofiber/fiber/v2"
	bookHandler "github.com/eld4niz/fiber-books-api/internal/handlers/book"
)

func SetupBookRoutes(router fiber.Router) {
	router.Get("/books", bookHandler.GetBooks)
	router.Get("/books/:bookId", bookHandler.GetBook)
	router.Post("/books", bookHandler.CreateBook)
	router.Put("/books/:bookId", bookHandler.UpdateBook)
	router.Delete("/books/:bookId", bookHandler.DeleteBook)
}