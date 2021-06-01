package repository

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"
)

type Book interface {
	GetBooks(context.Context) ([]models.Book, error)
	GetBook(context.Context, int64) (models.Book, error)
	DeleteBook(context.Context, int64) error
}

type Repository struct {
	Book
}

func NewRepository() *Repository {
	return &Repository{}
}
