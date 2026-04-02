package handlers

import (
	"fmt"
	"net/http"
	"strconv" // Necessário para converter string para int

	"github.com/felippevianna/go-api-livros/internal/models"
	"github.com/felippevianna/go-api-livros/internal/repository"
	service "github.com/felippevianna/go-api-livros/internal/service/googleBooks"
	"github.com/gin-gonic/gin"
)

type CreateBookRequest struct {
	Titulo       string `json:"titulo" binding:"required"`
	Descricao    string `json:"descricao"`
	AutorID      uint   `json:"autor_id" binding:"required"`
	CategoriaIDs []uint `json:"categoria_ids"`
}

type BookHandler struct {
	repo repository.LivroRepository
}

func NewBookHandler(repo repository.LivroRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	userId, _ := c.Get("userID")

	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	livro := models.Livro{
		Titulo:    req.Titulo,
		Descricao: req.Descricao,
		AutorID:   req.AutorID,
		UserID:    userId.(uint),
	}

	if err := h.repo.CreateWithCategories(&livro, req.CategoriaIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar livro e categorias"})
		return
	}

	c.JSON(http.StatusCreated, livro)
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	livros, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livros"})
		return
	}

	c.JSON(http.StatusOK, livros)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	// Pega o parâmetro :id da URL como string
	idStr := c.Param("id")

	// Converte string para uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chama o repositório
	livro, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Tenta deletar
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar livro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro removido com sucesso"})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var livro models.Livro
	// Pegamos os novos dados do corpo da requisição
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	livro.ID = uint(id) // Garantimos que o ID da struct seja o da URL

	if err := h.repo.Update(&livro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar livro"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func (h *BookHandler) SearchBooks(c *gin.Context) {
	// Captura os parâmetros da URL
	titulo := c.Query("titulo")
	autorIDStr := c.Query("autor_id")

	var autorID uint64
	if autorIDStr != "" {
		var err error
		autorID, err = strconv.ParseUint(autorIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID do autor inválido"})
			return
		}
	}

	// Chama o método do repositório
	livros, err := h.repo.Search(titulo, uint(autorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livros"})
		return
	}

	c.JSON(http.StatusOK, livros)
}

func (h *BookHandler) SearchBooksGoogleApi(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro de busca 'q' é obrigatório"})
		return
	}

	result, err := service.SearchGoogleBooks(query)
	// No seu BookHandler
	if err != nil {
		// Esse Println vai aparecer no terminal do seu Docker
		fmt.Println("ERRO REAL DA API DO GOOGLE:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Retorna o erro real pro Postman
		return
	}

	c.JSON(http.StatusOK, result)
}
