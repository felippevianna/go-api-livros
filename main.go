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
	authorRepo := repository.NewAuthorRepository(db)

	// 3. Instancia o handler passando o repositório (Injeção de Dependência)
	bookHandler := handlers.NewBookHandler(livroRepo)
	authorHandler := handlers.NewAuthorHandler(authorRepo)

	// 4. Configura o Gin
	r := gin.Default()

	// Rotas de Autores
	r.POST("/authors", authorHandler.CreateAuthor)
	r.GET("/authors", authorHandler.GetAuthors)

	// Rotas de Livros
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.DELETE("/books/:id", bookHandler.DeleteBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.GET("/books/search", bookHandler.SearchBooks)

	// Roda o servidor
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao rodar o servidor: %v", err)
	}
}