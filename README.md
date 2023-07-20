# Bookstore API

This is a RESTful API for a bookstore. It provides endpoints for creating, retrieving, updating, and deleting books. The API is built using the Go programming language and the Gorilla Mux router.

## Endpoints

- **POST /book/**: Create a new book
- **GET /books/**: Get all books
- **GET /book/{bookId}**: Get a book by ID
- **PUT /book/{bookId}**: Update a book by ID
- **DELETE /book/{bookId}**: Delete a book by ID

## Project Structure

The project is structured as follows:

- **pkg/controllers**: Contains the controller functions for handling API requests.
- **pkg/models**: Contains the data models for the bookstore API.
- **pkg/routes**: Contains the router configuration for the bookstore API.
- **main.go**: The entry point of the application.

## Dependencies

The project uses the following dependencies:

- **github.com/gorilla/mux**: A powerful HTTP router and URL matcher for building Go web servers.
- **github.com/jinzhu/gorm**: A powerful ORM library for Go that provides a simple and clean API for working with databases.
