package repository

import (
	"go-api-livros/internal/models"
	"gorm.io/gorm"
)

type AutorRepository interface {
	Create(autor *models.Autor) error
	FindAll() ([]models.Autor, error)
}

type autorRepository struct {
	db *gorm.DB
}

func NewAutorRepository(db *gorm.DB) AutorRepository {
	return &autorRepository{db: db}
}

func (r *autorRepository) Create(autor *models.Autor) error {
	return r.db.Create(autor).Error
}

func (r *autorRepository) FindAll() ([]models.Autor, error) {
	var autores []models.Autor
	// Preloads: Carrega os livros do autor na mesma consulta (Opcional)
	err := r.db.Preload("Livros").Find(&autores).Error
	return autores, err
}