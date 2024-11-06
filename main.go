package main

import (
	"log"
	"os"

	"github.com/AhmedRabea0302/fiber-postgres/models"
	"github.com/AhmedRabea0302/fiber-postgres/storage"
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
	books := &[]models.Books{}
	err := r.DB.Find(books).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Error fetching books",
		})
	}

	ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Books fetched successfully",
		"data":    books,
	})

	return nil
}

func (r Repository) CreateBook(ctx *fiber.Ctx) error {
	book := Book{}

	err := ctx.BodyParser(&book)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Error parsing JSON",
		})
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Couldn't Create Book",
		})
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Book has been created",
	})

	return nil
}

func (r Repository) DeleteBook(ctx *fiber.Ctx) error {
	book := models.Books{}
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Missing book ID",
		})
		return nil
	}

	err := r.DB.Delete(book, id)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Couldn't delete book",
		})
		return err.Error
	}

	ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Book has been deleted",
	})

	return nil
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  true, // change to false if you're not using SSL
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
