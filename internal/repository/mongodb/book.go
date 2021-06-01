package mongodb

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbBookRepository struct {
	db *mongo.Database
}

func (mr *MongodbBookRepository) GetBook(ctx context.Context, hexID string) (models.Book, error) {
	id, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return models.Book{}, err
	}

	filter := bson.M{"_id": id}
	var book models.Book
	err = mr.db.Collection("books").FindOne(ctx, filter).Decode(&book)
	return book, err
}

func (mr *MongodbBookRepository) DeleteBook(ctx context.Context, hexID string) error {
	id, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	_, err = mr.db.Collection("books").DeleteOne(ctx, filter)
	return err
}

func NewMongodbRepository(db *mongo.Database) *MongodbBookRepository {
	return &MongodbBookRepository{
		db: db,
	}
}
