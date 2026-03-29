package handlers

import (
	"net/http"
	"strconv"
	"github.com/felippevianna/go-api-livros/internal/models"
	"github.com/felippevianna/go-api-livros/internal/repository"
	"github.com/gin-gonic/gin"
)

type AvaliacaoHandler struct {
	repo repository.AvaliacaoRepository
}

func NewAvaliacaoHandler(repo repository.AvaliacaoRepository) *AvaliacaoHandler {
	return &AvaliacaoHandler{repo: repo}
}

func (h *AvaliacaoHandler) CreateReview(c *gin.Context) {
	var av models.Avaliacao
	if err := c.ShouldBindJSON(&av); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validação simples de nota via código
	if av.Nota < 1 || av.Nota > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A nota deve ser entre 1 e 5"})
		return
	}

	if err := h.repo.Create(&av); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar avaliação"})
		return
	}

	c.JSON(http.StatusCreated, av)
}

func (h *AvaliacaoHandler) GetReviewsByBook(c *gin.Context) {
	// Pega o ID que virá na URL: /books/:id/reviews
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do livro inválido"})
		return
	}

	reviews, err := h.repo.FindByLivroID(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar avaliações"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}