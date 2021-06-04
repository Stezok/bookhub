package api

import (
	"database/sql"
	"net/http"

	"github.com/Stezok/bookhub/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *APIHandler) CreateBookHandler(c *gin.Context) {
	var book JSONBook
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook := models.Book{
		Title: book.Title,
		Desc:  book.Desc,
	}

	id, err := h.service.CreateBook(c.Request.Context(), newBook)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookID": id})
}

func (h *APIHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.service.GetBooks(c.Request.Context())
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"result": "BookHub is empty"})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Print(err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *APIHandler) GetBookHandler(c *gin.Context) {
	var uriData UriBookID
	if err := c.BindUri(&uriData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.GetBook(c.Request.Context(), *uriData.ID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"result": "no matches with provided id"})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": book})
}

func (h *APIHandler) UpdateBookHandler(c *gin.Context) {
	var jsonData JSONBook
	if err := c.BindJSON(&jsonData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var uriData UriBookID
	if err := c.BindUri(&uriData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook := models.Book{
		ID:    *uriData.ID,
		Title: jsonData.Title,
		Desc:  jsonData.Desc,
	}

	newBook, err := h.service.UpdateBook(c.Request.Context(), updateBook)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Print(err)
		return
	}

	c.JSON(http.StatusOK, newBook)
}

func (h *APIHandler) DeleteBookHandler(c *gin.Context) {
	var uriData UriBookID
	if err := c.BindUri(&uriData); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Print(err)
		return
	}

	err := h.service.DeleteBook(c.Request.Context(), *uriData.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "deleted"})
}
