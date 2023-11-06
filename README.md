# Fiber Framework REST API

This is a RESTful API written in Go, utilizing the Fiber framework and connected to a PostgreSQL database.

## Prerequisites

Before running the application, ensure you have the following installed:

- Go (at least version 1.19)
- PostgreSQL
- [Fiber v2](https://github.com/gofiber/fiber)

## Installing Dependencies

To install the dependencies, run the following command:

```bash
go mod download
```

## Project Structure

The project has the following structure:

- `config/`: Contains configuration files.
    - `config.go`: Handles application configuration.
- `database/`: Handles database connections.
    - `connect.go`: Manages PostgreSQL database connection.
- `internal/`: Internal packages.
    - `handlers/`: Request handlers.
        - `book/`: Handles book-related requests.
            - `book.go`: Contains book-related handler functions.
    - `model/`: Data models.
        - `model.go`: Defines the data model.
    - `routes/`: Application routes.
        - `books/`: Defines routes for managing books.
            - `book.go`: Contains book-related routes.
- `main.go`: Main entry point of the application.
- `Makefile`: Contains useful commands for building and running the project.
- `router/`: Handles routing logic.
    - `router.go`: Defines application routes.
- `server`: Compiled binary of the application.

## Running the Application

To run the application, run the following command:

#### Ways to Run
```bash
make watch
```
OR
```bash
go run main.go
```
