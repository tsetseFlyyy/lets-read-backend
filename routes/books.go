package routes

import (
	"context"
	"fmt"
	"net/http"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var booksCollection *mongo.Collection = OpenCollection(Client, "books")

func AddBook(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(book)
	fmt.Println(book.Name)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	if book.Title == "" && book.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Author properties are empty"})
		return
	}

	book.ID = primitive.NewObjectID()

	result, insertErr := booksCollection.InsertOne(ctx, book)
	if insertErr != nil {
		msg := fmt.Sprintf("book item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusCreated, result)
}

func GetBooks(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var books []bson.M

	cursor, err := booksCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, books)
}

func UpdateBook(c *gin.Context) {

	bookID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(bookID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(book)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, err := booksCollection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$set", bson.M{"surname": book.Surname, "name": book.Name, "patronymic": book.Patronymic, "deadline": book.Deadline}},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusInternalServerError, result.ModifiedCount)
	} else if result.ModifiedCount == 1 {
		c.JSON(http.StatusOK, result.ModifiedCount)
	}
}

func UpdateNotes(c *gin.Context) {
	bookID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(bookID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var book models.Note

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	fmt.Println("not:", book.Not)

	validationErr := validate.Struct(book)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, err := booksCollection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$push", bson.M{"notes": book.Not}},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusInternalServerError, result.ModifiedCount)
	} else if result.ModifiedCount == 1 {
		c.JSON(http.StatusOK, result.ModifiedCount)
	}
}

func DeleteBook(c *gin.Context) {

	bookID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(bookID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := booksCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	if result.DeletedCount == 0 {
		c.JSON(http.StatusInternalServerError, result.DeletedCount)
	} else if result.DeletedCount == 1 {
		c.JSON(http.StatusOK, result.DeletedCount)
	}

}
