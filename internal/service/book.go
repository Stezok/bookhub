package service

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"
)

type BookRepo interface {
	GetBooks(context.Context) ([]models.Book, error)
	GetBook(context.Context, int64) (models.Book, error)
	DeleteBook(context.Context, int64) error
}

type BookService struct {
	repo BookRepo
}

func (bs *BookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	return bs.repo.GetBooks(ctx)
}

func (bs *BookService) GetBook(ctx context.Context, id int64) (models.Book, error) {
	return bs.repo.GetBook(ctx, id)
}

func (bs *BookService) DeleteBook(ctx context.Context, id int64) error {
	return bs.repo.DeleteBook(ctx, id)
}

func NewBookService(repo BookRepo) *BookService {
	return &BookService{
		repo: repo,
	}
}
