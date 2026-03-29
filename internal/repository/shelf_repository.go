package repository

import (
	"github.com/felippevianna/go-api-livros/internal/models"
	"gorm.io/gorm"
	"time"
)

type ShelfRepository interface {
	AddToShelf(shelf *models.Shelf) error
	GetByUserID(userID uint) ([]models.Shelf, error)
	UpdateStatus(userID, shelfID uint, status models.ReadingStatus) error
}

type shelfRepository struct {
	db *gorm.DB
}

func NewShelfRepository(db *gorm.DB) ShelfRepository {
	return &shelfRepository{db: db}
}

func (r *shelfRepository) AddToShelf(shelf *models.Shelf) error {
	return r.db.Create(shelf).Error
}

func (r *shelfRepository) GetByUserID(userID uint) ([]models.Shelf, error) {
	var items []models.Shelf
	err := r.db.Preload("Livro").Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

func (r *shelfRepository) UpdateStatus(userID, shelfID uint, status models.ReadingStatus) error {
	updates := map[string]interface{}{"status": status}
	
	// Se mudar para lido, atualizamos a data de finalização
	if status == models.StatusLido {
		updates["finished_at"] = time.Now()
	}

	return r.db.Model(&models.Shelf{}).
		Where("id = ? AND user_id = ?", shelfID, userID).
		Updates(updates).Error
}