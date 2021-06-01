package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	ShortDesc string             `bson:"short_desc"`
	Desc      string             `bson:"desc"`
}
