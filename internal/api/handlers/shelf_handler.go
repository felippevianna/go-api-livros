package handlers

import (
	"net/http"
	"time"
	"github.com/felippevianna/go-api-livros/internal/models"
	"github.com/felippevianna/go-api-livros/internal/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ShelfHandler struct {
	repo repository.ShelfRepository
}

func NewShelfHandler(repo repository.ShelfRepository) *ShelfHandler {
	return &ShelfHandler{repo: repo}
}

func (h *ShelfHandler) AddToShelf(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req struct {
		LivroID uint                 `json:"livro_id" binding:"required"`
		Status  models.ReadingStatus `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelf := models.Shelf{
		UserID:  userID.(uint),
		LivroID: req.LivroID,
		Status:  req.Status,
	}

	if req.Status == models.StatusLido {
		now := time.Now()
		shelf.FinishedAt = &now
	}

	if err := h.repo.AddToShelf(&shelf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar à estante"})
		return
	}

	c.JSON(http.StatusCreated, shelf)
}

func (h *ShelfHandler) GetMyShelf(c *gin.Context) {
	userID, _ := c.Get("userID")

	items, err := h.repo.GetByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar estante"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *ShelfHandler) UpdateShelfStatus(c *gin.Context) {
	userID, _ := c.Get("userID")
	shelfID, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Status models.ReadingStatus `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateStatus(userID.(uint), uint(shelfID), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar estante"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado!"})
}