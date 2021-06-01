package mysql

import (
	"context"

	"github.com/Stezok/bookhub/internal/models"

	"github.com/jmoiron/sqlx"
)

type MySQLBookRepository struct {
	db *sqlx.DB
}

func (br *MySQLBookRepository) GetBooks(ctx context.Context) ([]models.Book, error) {
	var book []models.Book
	query := `SELECT * FROM Book`
	err := br.db.SelectContext(ctx, &book, query)
	return book, err
}

func (br *MySQLBookRepository) GetBook(ctx context.Context, id int64) (models.Book, error) {
	var book models.Book
	query := `SELECT * FROM Book WHERE id = ?`
	err := br.db.GetContext(ctx, &book, query, id)
	return book, err
}

func (br *MySQLBookRepository) DeleteBook(ctx context.Context, id int64) error {
	query := `DELETE FROM Book WHERE id = ?`
	_, err := br.db.ExecContext(ctx, query, id)
	return err
}

func NewMySQLBookRepository(db *sqlx.DB) *MySQLBookRepository {
	return &MySQLBookRepository{
		db: db,
	}
}
