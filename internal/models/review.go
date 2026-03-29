package models

import (
	"time"
	"gorm.io/gorm"
)

type Avaliacao struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	LivroID    uint           `gorm:"not null" json:"livro_id"` 
	Nota       int            `gorm:"not null" json:"nota"`
	Comentario string         `gorm:"type:text" json:"comentario"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"` // O "-" esconde do JSON
}