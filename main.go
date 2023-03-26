package main

import (
	"net/http"
	"os"

	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	router.GET("", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Connection is ok!"})
	})
	router.GET("/books", routes.GetBooks)
	router.POST("/books", routes.AddBook)
	router.PUT("/books/:id", routes.UpdateBook)
	router.PUT("/books/notes/:id", routes.UpdateNotes)
	router.DELETE("/books/:id", routes.DeleteBook)
	router.Run(":" + port)
}
