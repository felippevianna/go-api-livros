package handlers

import (
	"net/http"
	"github.com/felippevianna/go-api-livros/internal/models"
	"github.com/felippevianna/go-api-livros/internal/repository"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	repo repository.AuthorRepository
}

func NewAuthorHandler(repo repository.AuthorRepository) *AuthorHandler {
	return &AuthorHandler{repo: repo}
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar autor"})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	autores, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar autores"})
		return
	}
	c.JSON(http.StatusOK, autores)
}