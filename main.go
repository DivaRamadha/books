package main

import (
	"fmt"
	"log"
	"main/controllers"
	"main/models"
	"main/repositories"
	"main/services"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Book{}, &models.Author{}, &models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

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

	router := controllers.NewRouter(bookController, authorController, userController, authController, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
