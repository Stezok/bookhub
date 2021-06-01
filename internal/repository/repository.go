package repository

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"
)

type Book interface {
	GetBook(context.Context, string) (models.Book, error)
	DeleteBook(context.Context, string) error
}

type Repository struct {
	Book
}

func NewRepository() *Repository {
	return &Repository{}
}
