package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Author      string             `bson:"author"`
	FriendsBook bool               `bson:"friendsbook"`
	Surname     string             `bson:"surname"`
	Name        string             `bson:"name"`
	Patronymic  string             `bson:"patronymic"`
	Deadline    string             `bson:"deadline"`
	Notes       []string           `bson:"notes,omitempty"`
}

type Note struct {
	Not string `bson:"not,omitempty"`
}
