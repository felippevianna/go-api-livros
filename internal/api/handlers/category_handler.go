package handlers

import (
	"net/http"
	"github.com/felippevianna/go-api-livros/internal/models"
	"github.com/felippevianna/go-api-livros/internal/repository"
	"github.com/gin-gonic/gin"
)

type CategoriaHandler struct {
	repo repository.CategoriaRepository
}

func NewCategoriaHandler(repo repository.CategoriaRepository) *CategoriaHandler {
	return &CategoriaHandler{repo: repo}
}

func (h *CategoriaHandler) CreateCategory(c *gin.Context) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&categoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar categoria"})
		return
	}

	c.JSON(http.StatusCreated, categoria)
}

func (h *CategoriaHandler) GetCategories(c *gin.Context) {
	categorias, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar categorias"})
		return
	}
	c.JSON(http.StatusOK, categorias)
}