package repository

import (
	"github.com/felippevianna/go-api-livros/internal/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	Create(author *models.Author) error
	FindAll() ([]models.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db: db}
}

func (r *authorRepository) Create(author *models.Author) error {
	return r.db.Create(author).Error
}

func (r *authorRepository) FindAll() ([]models.Author, error) {
	var autores []models.Author
	// Preloads: Carrega os livros do autor na mesma consulta (Opcional)
	err := r.db.Preload("Livros").Find(&autores).Error
	return autores, err
}