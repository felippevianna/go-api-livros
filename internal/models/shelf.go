package models

import (
	"time"
	"gorm.io/gorm"
)

type Shelf struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	LivroID   uint      `gorm:"not null" json:"livro_id"`
	Status    string    `gorm:"type:varchar(50);default:'quero_ler'" json:"status"` // lido, lendo, quero_ler
	Livro     Livro     `gorm:"foreignKey:LivroID" json:"livro"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}