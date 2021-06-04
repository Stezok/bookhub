package service

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"
)

type BookRepo interface {
	CreateBook(context.Context, models.Book) (int64, error)
	GetBooks(context.Context) ([]models.Book, error)
	GetBook(context.Context, int64) (models.Book, error)
	UpdateBook(context.Context, models.Book) error
	DeleteBook(context.Context, int64) error
}

type BookService struct {
	repo BookRepo
}

func (bs *BookService) CreateBook(ctx context.Context, book models.Book) (int64, error) {
	return bs.repo.CreateBook(ctx, book)
}

func (bs *BookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	return bs.repo.GetBooks(ctx)
}

func (bs *BookService) GetBook(ctx context.Context, id int64) (models.Book, error) {
	return bs.repo.GetBook(ctx, id)
}

func (bs *BookService) updateBook(oldBook, newBook models.Book) models.Book {
	newBook.ID = oldBook.ID
	if newBook.Desc == "" {
		newBook.Desc = oldBook.Desc
	}
	if newBook.Title == "" {
		newBook.Title = oldBook.Title
	}
	return newBook
}

func (bs *BookService) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {
	oldBook, err := bs.repo.GetBook(ctx, book.ID)
	if err != nil {
		return models.Book{}, err
	}

	book = bs.updateBook(oldBook, book)
	err = bs.repo.UpdateBook(ctx, book)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (bs *BookService) DeleteBook(ctx context.Context, id int64) error {
	return bs.repo.DeleteBook(ctx, id)
}

func NewBookService(repo BookRepo) *BookService {
	return &BookService{
		repo: repo,
	}
}
