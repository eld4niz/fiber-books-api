package bookHandler

import (
	"github.com/eld4niz/fiber-books-api/database"
	"github.com/eld4niz/fiber-books-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []model.Book

	// find all books in the database
	db.Find(&books)

	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No books provided", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Book Found", "data": books})
}

func CreateBook(c *fiber.Ctx) error {
	db := database.DB
	book := new(model.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	book.ID = uuid.New()

	err = db.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create book", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Book", "data": book})
}

func GetBook(c *fiber.Ctx) error {
	db := database.DB
	var book model.Book

	// Read the param bookId
	id := c.Params("bookId")

	// Find the book with the given ID
	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Book Found", "data": book})
}

func UpdateBook(c *fiber.Ctx) error {
	type updateBook struct {
		BookName   string  `json:"book_name"`
		AuthorName string  `json:"author_name"`
		Category   string  `json:"category"`
		Price      float64 `json:"price"`
	}
	db := database.DB
	var book model.Book

	// Read the param bookId
	id := c.Params("bookId")

	// Find the book with the given Id
	db.Find(&book, "id = ?", id)

	// If no such book present return an error
	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book provided", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateBookData updateBook
	err := c.BodyParser(&updateBookData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	book.BookName = updateBookData.BookName
	book.AuthorName = updateBookData.AuthorName
	book.Category = updateBookData.Category
	book.Price = updateBookData.Price

	// Save the Changes
	db.Save(&book)

	// Return the updated book
	return c.JSON(fiber.Map{"status": "success", "message": "Books Found", "data": book})
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DB
	var book model.Book

	// Read the param bookId
	id := c.Params("bookId")

	// Find the book with the given Id
	db.Find(&book, "id = ?", id)

	// If no such book present return an error
	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book found", "data": nil})
	}

	// Delete the book and return error if encountered
	err := db.Delete(&book, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete book", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Book"})
}
