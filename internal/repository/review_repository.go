package repository

import (
	"github.com/felippevianna/go-api-livros/internal/models"
	"gorm.io/gorm"
)

type AvaliacaoRepository interface {
	Create(avaliacao *models.Avaliacao) error
	FindByLivroID(livroID uint) ([]models.Avaliacao, error)
}

type avaliacaoRepository struct {
	db *gorm.DB
}

func NewAvaliacaoRepository(db *gorm.DB) AvaliacaoRepository {
	return &avaliacaoRepository{db: db}
}

func (r *avaliacaoRepository) Create(av *models.Avaliacao) error {
	return r.db.Create(av).Error
}

func (r *avaliacaoRepository) FindByLivroID(id uint) ([]models.Avaliacao, error) {
	var avs []models.Avaliacao
	err := r.db.Where("livro_id = ?", id).Find(&avs).Error
	return avs, err
}