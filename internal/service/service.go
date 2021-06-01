package service

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"
)

type Book interface {
	GetBook(context.Context, string) (models.Book, error)
	DeleteBook(context.Context, string) error
}

type Service struct {
	Book
}

func NewService() *Service {
	return &Service{}
}
