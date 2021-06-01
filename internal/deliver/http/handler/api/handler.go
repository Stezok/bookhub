package api

import (
	"log"

	"github.com/Stezok/bookhub/internal/service"
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	service service.Service
	logger  *log.Logger
}

func (h *APIHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/books", h.GetBooksHandler)

	router.GET("/book/:id", h.GetBookHandler)
	router.DELETE("/book/:id", h.DeleteBookHandler)

	return router
}

func NewAPIHandler(service service.Service, logger *log.Logger) *APIHandler {
	return &APIHandler{
		service: service,
		logger:  logger,
	}
}
