package handler

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/brunnossanttos/libri/internal/service"
)

type BookHandler struct {
  Svc *service.BookService
}

func NewBookHandler(s *service.BookService) *BookHandler {
  return &BookHandler{Svc: s}
}

func (h *BookHandler) SearchBooks(c *gin.Context) {
  q := c.Query("q")
  if q == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'q' é obrigatório"})
    return
  }

  books, err := h.Svc.SearchBooks(c.Request.Context(), q)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, books)
}
