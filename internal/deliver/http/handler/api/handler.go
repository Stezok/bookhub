package api

import (
	"context"
	"log"

	"github.com/Stezok/bookhub/internal/models"
	"github.com/gin-gonic/gin"
)

type BookService interface {
	CreateBook(context.Context, models.Book) (int64, error)
	GetBooks(context.Context) ([]models.Book, error)
	GetBook(context.Context, int64) (models.Book, error)
	UpdateBook(context.Context, models.Book) (models.Book, error)
	DeleteBook(context.Context, int64) error
}

type Service interface {
	BookService
}

type APIHandler struct {
	service Service
	logger  *log.Logger
}

func (h *APIHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/new/book", h.CreateBookHandler)

	router.GET("/books", h.GetBooksHandler)

	router.PUT("/book/:id", h.UpdateBookHandler)
	router.GET("/book/:id", h.GetBookHandler)
	router.DELETE("/book/:id", h.DeleteBookHandler)

	return router
}

func NewAPIHandler(service Service, logger *log.Logger) *APIHandler {
	return &APIHandler{
		service: service,
		logger:  logger,
	}
}
