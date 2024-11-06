package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) SetupRputes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("book/create", r.CreateBook)
	api.Delete("book/delete", r.DeleteBook)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRputes(app)

	app.Listen(":8080")

}
