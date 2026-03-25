package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/felippevianna/go-api-livros/config"
	"github.com/felippevianna/go-api-livros/internal/api/handlers"
	"github.com/felippevianna/go-api-livros/internal/repository" // Importação necessária
)

func main() {
	// 1. Conecta ao banco
	db := config.SetupDatabase()

	// 2. Instancia o repositório
	livroRepo := repository.NewLivroRepository(db)

	// 3. Instancia o handler passando o repositório (Injeção de Dependência)
	bookHandler := handlers.NewBookHandler(livroRepo)

	// 4. Configura o Gin
	r := gin.Default()

	// 5. Define as rotas
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	// Roda o servidor
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao rodar o servidor: %v", err)
	}
}