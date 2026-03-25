package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"github.com/felippevianna/go-api-livros/internal/models"
	"gorm.io/gorm"
)

// SetupDatabase configura a conexão com o banco de dados
func SetupDatabase() *gorm.DB {
	// Acessa as variáveis de ambiente que o Docker Compose injetou
	connectionSts := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	dbConnection, err := gorm.Open(postgres.Open(connectionSts), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados com GORM: %v", err)
	}

	fmt.Println("Conexão com o banco de dados GORM estabelecida com sucesso!")

	// Migra o esquema do banco de dados para criar a tabela de Livros
	// Isso é fundamental para que o GORM crie a tabela baseada na sua struct
	err = dbConnection.AutoMigrate(&models.Livro{}, &models.Autor{})

	if err != nil {
		log.Fatalf("Erro ao migrar o esquema do banco de dados: %v", err)
	}

	fmt.Println("Tabela de livros migrada com sucesso!")

	return dbConnection
}