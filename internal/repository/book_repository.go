package repository

import (
	"github.com/felippevianna/go-api-livros/internal/models"
	"gorm.io/gorm"
)

// LivroRepository é a interface que define os métodos do repositório de livros.
// Isso é uma boa prática para tornar o código mais testável.
type LivroRepository interface {
	Create(livro *models.Livro) error
	FindAll() ([]models.Livro, error)
	FindByID(id uint) (*models.Livro, error)
	Update(livro *models.Livro) error
	Delete(id uint) error
}

// livroRepository é a implementação da interface LivroRepository.
type livroRepository struct {
	db *gorm.DB
}

// NewLivroRepository cria uma nova instância do repositório de livros.
func NewLivroRepository(db *gorm.DB) LivroRepository {
	return &livroRepository{db: db}
}

func (r *livroRepository) Create(livro *models.Livro) error {
	return r.db.Create(livro).Error
}

func (r *livroRepository) FindAll() ([]models.Livro, error) {
	var livros []models.Livro
	// O nome dentro do Preload deve ser EXATAMENTE o nome do campo na struct Livro
	err := r.db.Preload("Autor").Find(&livros).Error
	return livros, err
}

func (r *livroRepository) FindByID(id uint) (*models.Livro, error) {
	var livro models.Livro
	// Buscam o autor junto com o livro específico
	err := r.db.Preload("Autor").First(&livro, id).Error
	if err != nil {
		return nil, err
	}
	return &livro, nil
}

func (r *livroRepository) Update(livro *models.Livro) error {
	return r.db.Save(livro).Error
}

func (r *livroRepository) Delete(id uint) error {
	return r.db.Delete(&models.Livro{}, id).Error
}