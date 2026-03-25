package handlers

import (
	"net/http"
	"github.com/felippevianna/go-api-livros/internal/models"
	"github.com/felippevianna/go-api-livros/internal/repository"
	"github.com/gin-gonic/gin"
	"strconv" // Necessário para converter string para int
)

type BookHandler struct {
	repo repository.LivroRepository
}

// NewBookHandler cria uma nova instância do handler injetando o repositório
func NewBookHandler(repo repository.LivroRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

// CreateBook lida com a requisição POST para criar um livro
func (h *BookHandler) CreateBook(c *gin.Context) {
	var livro models.Livro

	// Faz o "Bind" do JSON que vem no corpo da requisição para a struct Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o repositório para salvar no banco de dados
	if err := h.repo.Create(&livro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar livro"})
		return
	}

	// Retorna o livro criado com status 201 (Created)
	c.JSON(http.StatusCreated, livro)
}

// GetBooks lida com a requisição GET para listar todos os livros
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