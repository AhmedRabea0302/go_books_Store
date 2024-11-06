package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r Repository) SetupRputes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
	api.Post("book/", r.CreateBook)
	api.Delete("book/:id", r.DeleteBook)

}

// Get Book By ID
func (r Repository) GetBookByID(ctx *fiber.Ctx) error {
	return nil
}

func (r Repository) GetBooks(ctx *fiber.Ctx) error {
	return nil

}

func (r Repository) CreateBook(ctx *fiber.Ctx) error {
	return nil

}

func (r Repository) DeleteBook(ctx *fiber.Ctx) error {
	return nil

}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Couldn't connect to DB", err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRputes(app)

	app.Listen(":8080")

}
