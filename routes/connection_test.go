package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	//"go.mongodb.org/mongo-driver/mongo"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Author      string             `bson:"author"`
	FriendsBook bool               `bson:"friendsbook"`
	Surname     string             `bson:"surname"`
	Name        string             `bson:"name"`
	Patronymic  string             `bson:"patronymic"`
	Deadline    string             `bson:"deadline"`
	Notes       []string           `bson:"notes,omitempty"`
}

func TestHandler_getBooks(t *testing.T) {
	r := SetUpRouter()
	r.GET("/books", GetBooks)
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var books []Book
	json.Unmarshal(w.Body.Bytes(), &books)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)
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
