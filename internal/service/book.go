package service

import (
	"context"

	"github.com/Stezok/bookhub/internal/repository"

	"github.com/Stezok/bookhub/internal/models"
)

type BookService struct {
	repo repository.Book
}

func (bs *BookService) GetBook(ctx context.Context, hexID string) (models.Book, error) {
	return bs.repo.GetBook(ctx, hexID)
}

func (bs *BookService) DeleteBook(ctx context.Context, hexID string) error {
	return bs.repo.DeleteBook(ctx, hexID)
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{
		repo: repo,
	}
}
