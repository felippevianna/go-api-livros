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
	Search(titulo string, autorID uint) ([]models.Livro, error)
	CreateWithCategories(livro *models.Livro, categoriaIDs []uint) error
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
	err := r.db.Preload("Autor").Preload("Categorias").Find(&livros).Error
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

func (r *livroRepository) Delete(id uint) error {
	return r.db.Delete(&models.Livro{}, id).Error
}

func (r *livroRepository) Update(livro *models.Livro) error {
	// O Save atualiza todos os campos do modelo baseado no ID
	// Se preferir atualizar apenas campos preenchidos, usamos .Updates()
	return r.db.Model(&models.Livro{}).Where("id = ?", livro.ID).Updates(livro).Error
}

func (r *livroRepository) Search(titulo string, autorID uint) ([]models.Livro, error) {
    var livros []models.Livro
    query := r.db.Preload("Autor") // Já começamos com o Preload para trazer o autor

    // Se o título não estiver vazio, adiciona filtro ILIKE (case-insensitive no Postgres)
    if titulo != "" {
        // O % significa "qualquer coisa antes ou depois"
        query = query.Where("titulo ILIKE ?", "%"+titulo+"%")
    }

    // Se o autorID for maior que zero, filtra por ele
    if autorID > 0 {
        query = query.Where("autor_id = ?", autorID)
    }

    // Executa a query final
    err := query.Find(&livros).Error
    return livros, err
}

func (r *livroRepository) CreateWithCategories(livro *models.Livro, categoriaIDs []uint) error {
	// Se houver IDs de categorias, o GORM busca e associa automaticamente
	if len(categoriaIDs) > 0 {
		var categorias []models.Categoria
		if err := r.db.Find(&categorias, categoriaIDs).Error; err != nil {
			return err
		}
		livro.Categorias = categorias
	}
	return r.db.Create(livro).Error
}