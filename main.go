package main

import (
	"log"
	"main/controllers"
	"main/models"
	"main/repositories"
	"main/services"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/gobayarind?sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate models
	if err := db.AutoMigrate(&models.Book{}, &models.Author{}, &models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository, service, and controller
	bookRepository := repositories.NewBookRepository(db)
	authorRepository := repositories.NewAuthorRepository(db)
	userRepository := repositories.NewUserRepository(db)

	bookService := services.NewBookService(bookRepository)
	authorService := services.NewAuthorService(authorRepository)
	userService := services.NewUserService(userRepository)

	bookController := controllers.NewBookController(bookService)
	authorController := controllers.NewAuthorController(authorService)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(db)

	// Initialize the router and set up routes

	// Public routes

	// Initialize the router and set up routes
	router := controllers.NewRouter(bookController, authorController, userController, authController, db)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
