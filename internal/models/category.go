package models

import (
	"time"
	"gorm.io/gorm"
)

type Categoria struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Nome      string         `gorm:"type:varchar(100);unique;not null"`
	// Relacionamento Many-to-Many com Livros
	Livros    []Livro        `gorm:"many2many:livro_categorias;"` 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}