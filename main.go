package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shashank-kakarla/BookAPI/controllers"
	"github.com/shashank-kakarla/BookAPI/models"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()

	models.ConnectDatabase(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	router.GET("/books", controllers.FetchBooks)
	router.GET("/books/:id", controllers.FindBookByID)
	router.GET("/books/title/:title", controllers.FindBookByTitle)
	router.GET("/books/author/:author", controllers.FindBookByAuthor)

	router.POST("/books", controllers.CreateBook)

	router.PATCH("/books/:id", controllers.UpdateBook)

	router.DELETE("/books/:id", controllers.RemoveBook)

	router.Run()
}
