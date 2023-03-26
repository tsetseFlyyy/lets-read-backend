package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHandler_getBooks(t *testing.T) {
	r := SetUpRouter()
	r.GET("/books", GetBooks)
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var books []models.Book
	json.Unmarshal(w.Body.Bytes(), &books)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)
}

func TestHandler_addBook(t *testing.T) {
	r := SetUpRouter()
	r.POST("/books", AddBook)
	book := models.Book{
		Title: "book from test1",
		Author: "test",
		FriendsBook: false,
		
	}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandler_deleteBook(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/books/:id")
}

//func TestDBInstance(t *testing.T) {
//	client := DBinstance()

//	if client == nil {
//		t.Error("client is not existed")
//	}
//}

//var client *mongo.Client = DBinstance()

//func TestOpenCollection(t *testing.T) {
//	collection := OpenCollection(Client, "dwadadwa")

//	if collection == nil {
//		t.Error("collection is not existed")
//	}
//}
