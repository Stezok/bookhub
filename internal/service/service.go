package service

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"
)

type Book interface {
	CreateBook(context.Context, models.Book) (int64, error)
	GetBooks(context.Context) ([]models.Book, error)
	GetBook(context.Context, int64) (models.Book, error)
	UpdateBook(context.Context, models.Book) (models.Book, error)
	DeleteBook(context.Context, int64) error
}

type Service struct {
	Book
}

func NewService() *Service {
	return &Service{}
}
