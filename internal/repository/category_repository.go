package repository

import (
	"github.com/felippevianna/go-api-livros/internal/models"
	"gorm.io/gorm"
)

type CategoriaRepository interface {
	Create(categoria *models.Categoria) error
	FindAll() ([]models.Categoria, error)
	FindByID(id uint) (*models.Categoria, error)
}

type categoriaRepository struct {
	db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) CategoriaRepository {
	return &categoriaRepository{db: db}
}

func (r *categoriaRepository) Create(cat *models.Categoria) error {
	return r.db.Create(cat).Error
}

func (r *categoriaRepository) FindAll() ([]models.Categoria, error) {
	var categorias []models.Categoria
	return categorias, r.db.Find(&categorias).Error
}

func (r *categoriaRepository) FindByID(id uint) (*models.Categoria, error) {
	var cat models.Categoria
	err := r.db.First(&cat, id).Error
	return &cat, err
}