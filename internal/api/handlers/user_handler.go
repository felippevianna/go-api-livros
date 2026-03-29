package handlers

import (
	"net/http"
	"github.com/felippevianna/go-api-livros/internal/models"
	"github.com/felippevianna/go-api-livros/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserHandler struct {
	repo repository.UserRepository
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Senha string `json:"senha" binding:"required"`
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// O método Create do repositório já faz o Hash da senha automaticamente
	if err := h.repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
		return
	}

	// Retorna o usuário criado (a senha não aparecerá por causa do json:"-")
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Buscar usuário pelo email no repositório
	user, err := h.repo.FindByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "E-mail ou senha inválidos"})
		return
	}

	// DEBUG: Veja se o hash do banco e a senha da requisição chegaram
    fmt.Printf("Senha do banco (hash): %s\n", user.Senha)
    fmt.Printf("Senha digitada: %s\n", req.Senha)

	// 2. Comparar a senha digitada com o Hash do banco
	err = bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(req.Senha))
	if err != nil {
		// Se der erro, a senha está errada
		c.JSON(http.StatusUnauthorized, gin.H{"error": "E-mail ou senha inválidos"})
		return
	}

	// 3. Por enquanto, apenas retorna sucesso. 
	c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso!", "user_id": user.ID})
}