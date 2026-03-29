package models

import (
	"time"
	"gorm.io/gorm"
)

type ReadingStatus string

// Tipos de status para a estante do usuário. Isso ajuda a evitar erros de digitação e mantém a consistência.
const (
	StatusQueroLer ReadingStatus = "quero_ler"
	StatusLendo    ReadingStatus = "lendo"
	StatusLido     ReadingStatus = "lido"
	StatusAbandono ReadingStatus = "abandonado"
)

type Shelf struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  uint `gorm:"uniqueIndex:idx_user_livro" json:"user_id"`
	LivroID uint `gorm:"uniqueIndex:idx_user_livro" json:"livro_id"`

	// tipo ReadingStatus para garantir consistência
	Status    ReadingStatus  `gorm:"type:varchar(50);default:'quero_ler'" json:"status"`

	Livro     Livro     `gorm:"foreignKey:LivroID" json:"livro"`

	CreatedAt time.Time `json:"created_at"`
	FinishedAt *time.Time    `json:"finished_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}