package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/yeguacelestial/go-playground/fiber-rest-api/book"
	"github.com/yeguacelestial/go-playground/fiber-rest-api/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened.")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}
