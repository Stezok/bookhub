package service

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"
)

type Book interface {
	GetBooks(context.Context) ([]models.Book, error)
	GetBook(context.Context, int64) (models.Book, error)
	DeleteBook(context.Context, int64) error
}

type Service struct {
	Book
}

func NewService() *Service {
	return &Service{}
}
