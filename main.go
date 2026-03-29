package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/felippevianna/go-api-livros/config"
	"github.com/felippevianna/go-api-livros/internal/api/handlers"
	"github.com/felippevianna/go-api-livros/internal/api/middleware"
	"github.com/felippevianna/go-api-livros/internal/repository"
)

func main() {
	// 1. Conecta ao banco
	db := config.SetupDatabase()

	// 2. Instancia o repositório
	livroRepo := repository.NewLivroRepository(db)
	authorRepo := repository.NewAuthorRepository(db)
	categoriaRepo := repository.NewCategoriaRepository(db)
	avaliacaoRepo := repository.NewAvaliacaoRepository(db)
	userRepo := repository.NewUserRepository(db)
	shelfRepo := repository.NewShelfRepository(db)
	
	// 3. Instancia o handler passando o repositório (Injeção de Dependência)
	bookHandler := handlers.NewBookHandler(livroRepo)
	authorHandler := handlers.NewAuthorHandler(authorRepo)
	categoryHandler := handlers.NewCategoriaHandler(categoriaRepo)
	reviewHandler := handlers.NewAvaliacaoHandler(avaliacaoRepo)
	userHandler := handlers.NewUserHandler(userRepo)
	shelfHandler := handlers.NewShelfHandler(shelfRepo)


	// 4. Configura o Gin
	r := gin.Default()
	
	r.GET("/books", bookHandler.GetBooks)	
	
	// Rotas de Usuários
	r.POST("/users", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Rotas de Categorias
		protected.POST("/categories", categoryHandler.CreateCategory)
		protected.GET("/categories", categoryHandler.GetCategories)

		// Rotas de Autores
		protected.POST("/authors", authorHandler.CreateAuthor)
		protected.GET("/authors", authorHandler.GetAuthors)

		// Rotas de Livros
		protected.POST("/books", bookHandler.CreateBook)
		protected.GET("/books/:id", bookHandler.GetBookByID)
		protected.DELETE("/books/:id", bookHandler.DeleteBook)
		protected.PUT("/books/:id", bookHandler.UpdateBook)
		protected.GET("/books/search", bookHandler.SearchBooks)

		// Rotas de Avaliações
		protected.POST("/reviews", reviewHandler.CreateReview)
		protected.GET("/books/:id/reviews", reviewHandler.GetReviewsByBook)

		// Rotas de Estante
		protected.POST("/my-shelf", shelfHandler.AddToShelf)
   		protected.GET("/my-shelf", shelfHandler.GetMyShelf)
		// :id aqui é o ID do registro na tabela 'shelves', não do livro
		protected.PATCH("/my-shelf/:id", shelfHandler.UpdateShelfStatus)
		
	}

	// Roda o servidor
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao rodar o servidor: %v", err)
	}
}