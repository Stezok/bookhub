package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var uriData URIBookID
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

func (h *APIHandler) DeleteBookHandler(c *gin.Context) {
	var uriData URIBookID
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
