package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deshmukhpurushothaman/golang-snippets/go_postgres/models"
	"github.com/deshmukhpurushothaman/golang-snippets/go_postgres/storage"
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

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book) // converts the data from req to book type
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "create transaction failed"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book has been created!"})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the books"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "books fetched",
			"data": bookModels,
		})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not delete book"})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Deleted successfully!"})
	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := &models.Books{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	err := r.DB.Where("id=?", id).First(bookModel)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not find book"})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book found!", "data": bookModel})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete/:id", r.DeleteBook)
	api.Get("/get/:id", r.GetBookById)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Database connection failed")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("DB Migration failed!")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
